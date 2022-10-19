// msgq.go

package msgq

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// subscriber information
type subscriber struct {
	name string    // name is to identify a subscriber
	ix   int64     // the last index that's sent to this suscriber
	e    chan bool // set if a new data is posted
	done bool      // set when unsubscribed
	f    func(data interface{}) error
}

// MsgQ implements a simple pubsub system
type MsgQ struct {
	lock        *sync.RWMutex
	min, max    int // minimum and maximum number of entries
	ix2data     map[int64]interface{}
	head, tail  int64
	subscribers map[string]*subscriber
}

// NewMsgQ creates a new Msgq
func NewMsgQ(min, max int) *MsgQ {
	return &MsgQ{
		lock:        &sync.RWMutex{},
		min:         min,
		max:         max,
		ix2data:     map[int64]interface{}{},
		head:        0,
		tail:        -1,
		subscribers: map[string]*subscriber{},
	}
}

// Destroy is no-op for now
func (q *MsgQ) Destroy() {
}

// CountOfSubscribers returns the number of current subscribers
func (q *MsgQ) CountOfSubscribers() int {
	q.lock.RLock()
	count := len(q.subscribers)
	defer q.lock.RUnlock()
	return count
}

// Post adds a new data
func (q *MsgQ) Post(data interface{}) {
	q.lock.Lock()
	n := q.tail - q.head + 1
	ix := q.tail + 1
	q.ix2data[ix] = data
	q.tail = ix
	if n <= 0 {
		q.head = ix
	}
	n += 1
	bNeedTrimming := false
	if n%100 == 0 && n > int64(q.min) {
		bNeedTrimming = true
	}
	q.lock.Unlock()

	q.lock.RLock()
	for _, s := range q.subscribers {
		select {
		case s.e <- true:
			break
		default:
			break
		}
	}
	q.lock.RUnlock()

	if bNeedTrimming {
		q.Trim()
	}
}

// Trim cleans up old messages
func (q *MsgQ) Trim() {
	ix := int64(-1)

	q.lock.RLock()
	n := q.tail - q.head + 1
	max := q.tail - int64(q.max)

	if n <= int64(q.min) {
		q.lock.RUnlock()
		return
	}

	for _, s := range q.subscribers {
		if ix == -1 || s.ix < ix {
			ix = s.ix
		}
	}

	pass := (ix == -1 || ix == q.head) && ix >= max
	q.lock.RUnlock()
	if pass {
		return
	}

	q.lock.Lock()
	if ix < max {
		ix = max
	}
	if ix >= q.head {
		for i := q.head; i <= ix; i++ {
			delete(q.ix2data, i)
		}
		q.head = ix + 1
	}
	for _, s := range q.subscribers {
		if s.ix != -1 && s.ix < ix {
			s.ix = ix
		}
	}
	q.lock.Unlock()
}

// Subscribe adds a new subscriber
func (q *MsgQ) Subscribe(name string, f func(data interface{}) error) error {
	q.lock.Lock()
	if _, ok := q.subscribers[name]; ok {
		q.lock.Unlock()
		return errors.New("Already Exists")
	}

	s := &subscriber{
		name: name,
		ix:   -1,
		e:    make(chan bool, 1),
		done: false,
		f:    f,
	}
	q.subscribers[name] = s

	go func(q *MsgQ, s *subscriber) {
		for !s.done {
			<-s.e
			fmt.Printf("       !!! got event\n")

			for {
				six := s.ix + 1
				if six < q.head {
					six = q.head
				}
				eix := q.tail

				if six > eix {
					break
				}

				for i := six; i <= eix; i++ {
					var (
						d  interface{}
						ok bool
					)
					q.lock.RLock()
					d, ok = q.ix2data[i]
					q.lock.RUnlock()
					if !ok {
						continue
					} else {
						e := s.f(d)
						if e != nil {
							s.done = true
							break
						}
					}
				}
				s.ix = eix
			}
		}
	}(q, s)

	q.lock.Unlock()

	return nil
}

// Unsubscribe remoted the named subscriber
func (q *MsgQ) Unsubscribe(name string) error {
	var err error

	q.lock.Lock()
	if s, ok := q.subscribers[name]; !ok {
		err = errors.New("Not Found")
	} else {
		s.done = true
		select {
		case s.e <- true:
			break
		default:
			break
		}
	}
	delete(q.subscribers, name)
	q.lock.Unlock()
	return err
}

// msgq test
//
//nolint:deadcode
func msgqTest() error {
	q := NewMsgQ(10, 100)

	for x := 1; x <= 100; x++ {
		name := fmt.Sprintf("%d", x)
		q.Subscribe(name, func(data interface{}) error {
			if s, ok := data.(string); ok {
				fmt.Printf("%s: %s\n", name, s)
			}
			return nil
		})
	}

	i := 0
	for ; i < 50; i++ {
		q.Post(fmt.Sprintf("%d", i))
	}

	time.Sleep(2 * time.Second)

	for ; i < 100; i++ {
		q.Post(fmt.Sprintf("%d", i))
	}

	time.Sleep(2 * time.Second)

	for x := 2; x <= 100; x++ {
		name := fmt.Sprintf("%d", x)
		q.Unsubscribe(name)
	}

	time.Sleep(2 * time.Second)

	for ; i < 200; i++ {
		q.Post(fmt.Sprintf("%d", i))
	}

	time.Sleep(10 * time.Second)

	for x := 1; x <= 1; x++ {
		name := fmt.Sprintf("%d", x)
		q.Unsubscribe(name)
	}

	for {
		if q.CountOfSubscribers() == 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

// EOF
