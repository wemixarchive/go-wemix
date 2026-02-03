package influxdb

import (
	"fmt"
	uurl "net/url"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	client "github.com/influxdata/influxdb1-client/v2"
)

type reporter struct {
	reg      metrics.Registry
	interval time.Duration

	url       uurl.URL
	database  string
	username  string
	password  string
	namespace string
	tags      map[string]string

	client client.Client

	cache map[string]int64
}

// InfluxDB starts a InfluxDB reporter which will post the from the given metrics.Registry at each d interval.
func InfluxDB(r metrics.Registry, d time.Duration, url, database, username, password, namespace string) {
	InfluxDBWithTags(r, d, url, database, username, password, namespace, nil)
}

// InfluxDBWithTags starts a InfluxDB reporter which will post the from the given metrics.Registry at each d interval with the specified tags
func InfluxDBWithTags(r metrics.Registry, d time.Duration, url, database, username, password, namespace string, tags map[string]string) {
	u, err := uurl.Parse(url)
	if err != nil {
		log.Warn("Unable to parse InfluxDB", "url", url, "err", err)
		return
	}

	rep := &reporter{
		reg:       r,
		interval:  d,
		url:       *u,
		database:  database,
		username:  username,
		password:  password,
		namespace: namespace,
		tags:      tags,
		cache:     make(map[string]int64),
	}
	if err := rep.makeClient(); err != nil {
		log.Warn("Unable to make InfluxDB client", "err", err)
		return
	}

	rep.run()
}

// InfluxDBWithTagsOnce runs once an InfluxDB reporter and post the given metrics.Registry with the specified tags
func InfluxDBWithTagsOnce(r metrics.Registry, url, database, username, password, namespace string, tags map[string]string) error {
	u, err := uurl.Parse(url)
	if err != nil {
		return fmt.Errorf("unable to parse InfluxDB. url: %s, err: %v", url, err)
	}

	rep := &reporter{
		reg:       r,
		url:       *u,
		database:  database,
		username:  username,
		password:  password,
		namespace: namespace,
		tags:      tags,
		cache:     make(map[string]int64),
	}
	if err := rep.makeClient(); err != nil {
		return fmt.Errorf("unable to make InfluxDB client. err: %v", err)
	}

	if err := rep.send(); err != nil {
		return fmt.Errorf("unable to send to InfluxDB. err: %v", err)
	}

	return nil
}

func (r *reporter) makeClient() (err error) {
	r.client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     r.url.String(),
		Username: r.username,
		Password: r.password,
		Timeout:  10 * time.Second,
	})

	return
}

func (r *reporter) run() {
	intervalTicker := time.Tick(r.interval)
	pingTicker := time.Tick(time.Second * 5)

	for {
		select {
		case <-intervalTicker:
			if err := r.send(); err != nil {
				log.Warn("Unable to send to InfluxDB", "err", err)
			}
		case <-pingTicker:
			_, _, err := r.client.Ping(0)
			if err != nil {
				log.Warn("Got error while sending a ping to InfluxDB, trying to recreate client", "err", err)

				if err = r.makeClient(); err != nil {
					log.Warn("Unable to make InfluxDB client", "err", err)
				}
			}
		}
	}
}

func (r *reporter) send() error {
	var pts []*client.Point

	r.reg.Each(func(name string, i interface{}) {
		now := time.Now()
		namespace := r.namespace

		switch metric := i.(type) {
		case metrics.Counter:
			count := metric.Count()
			pt, err := client.NewPoint(
				fmt.Sprintf("%s%s.count", namespace, name),
				r.tags,
				map[string]interface{}{
					"value": count,
				},
				now)
			if err != nil {
				log.Warn("Unable to create InfluxDB point for counter", "name", name, "err", err)
				return
			}
			pts = append(pts, pt)
		case metrics.Gauge:
			ms := metric.Snapshot()
			pt, err := client.NewPoint(
				fmt.Sprintf("%s%s.gauge", namespace, name),
				r.tags,
				map[string]interface{}{
					"value": ms.Value(),
				},
				now)
			if err != nil {
				log.Warn("Unable to create InfluxDB point for gauge", "name", name, "err", err)
				return
			}
			pts = append(pts, pt)
		case metrics.GaugeFloat64:
			ms := metric.Snapshot()
			pt, err := client.NewPoint(
				fmt.Sprintf("%s%s.gauge", namespace, name),
				r.tags,
				map[string]interface{}{
					"value": ms.Value(),
				},
				now)
			if err != nil {
				log.Warn("Unable to create InfluxDB point for gauge float64", "name", name, "err", err)
				return
			}
			pts = append(pts, pt)
		case metrics.Histogram:
			ms := metric.Snapshot()
			if ms.Count() > 0 {
				ps := ms.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999, 0.9999})

				pt, err := client.NewPoint(
					fmt.Sprintf("%s%s.histogram", namespace, name),
					r.tags,
					map[string]interface{}{
						"count":    ms.Count(),
						"max":      ms.Max(),
						"mean":     ms.Mean(),
						"min":      ms.Min(),
						"stddev":   ms.StdDev(),
						"variance": ms.Variance(),
						"p50":      ps[0],
						"p75":      ps[1],
						"p95":      ps[2],
						"p99":      ps[3],
						"p999":     ps[4],
						"p9999":    ps[5],
					},
					now)
				if err != nil {
					log.Warn("Unable to create InfluxDB point for histogram", "name", name, "err", err)
					return
				}
				pts = append(pts, pt)
			}
		case metrics.Meter:
			ms := metric.Snapshot()
			pt, err := client.NewPoint(
				fmt.Sprintf("%s%s.meter", namespace, name),
				r.tags,
				map[string]interface{}{
					"count": ms.Count(),
					"m1":    ms.Rate1(),
					"m5":    ms.Rate5(),
					"m15":   ms.Rate15(),
					"mean":  ms.RateMean(),
				},
				now)
			if err != nil {
				log.Warn("Unable to create InfluxDB point for meter", "name", name, "err", err)
				return
			}
			pts = append(pts, pt)
		case metrics.Timer:
			ms := metric.Snapshot()
			ps := ms.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999, 0.9999})
			pt, err := client.NewPoint(
				fmt.Sprintf("%s%s.timer", namespace, name),
				r.tags,
				map[string]interface{}{
					"count":    ms.Count(),
					"max":      ms.Max(),
					"mean":     ms.Mean(),
					"min":      ms.Min(),
					"stddev":   ms.StdDev(),
					"variance": ms.Variance(),
					"p50":      ps[0],
					"p75":      ps[1],
					"p95":      ps[2],
					"p99":      ps[3],
					"p999":     ps[4],
					"p9999":    ps[5],
					"m1":       ms.Rate1(),
					"m5":       ms.Rate5(),
					"m15":      ms.Rate15(),
					"meanrate": ms.RateMean(),
				},
				now)
			if err != nil {
				log.Warn("Unable to create InfluxDB point for timer", "name", name, "err", err)
				return
			}
			pts = append(pts, pt)
		case metrics.ResettingTimer:
			t := metric.Snapshot()

			if len(t.Values()) > 0 {
				ps := t.Percentiles([]float64{50, 95, 99})
				val := t.Values()
				pt, err := client.NewPoint(
					fmt.Sprintf("%s%s.span", namespace, name),
					r.tags,
					map[string]interface{}{
						"count": len(val),
						"max":   val[len(val)-1],
						"mean":  t.Mean(),
						"min":   val[0],
						"p50":   ps[0],
						"p95":   ps[1],
						"p99":   ps[2],
					},
					now)
				if err != nil {
					log.Warn("Unable to create InfluxDB point for resetting timer", "name", name, "err", err)
					return
				}
				pts = append(pts, pt)
			}
		}
	})

	bps, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: r.database,
	})
	if err != nil {
		return err
	}
	bps.AddPoints(pts)
	err = r.client.Write(bps)
	return err
}
