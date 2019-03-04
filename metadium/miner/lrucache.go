// lrucache.go

package miner

import (
	"container/list"
	"sync"
)

type LruCache struct {
	lock  *sync.RWMutex
	max   int
	count int
	lru   *list.List
	data  map[interface{}]interface{}
}

func NewLruCache(max int) *LruCache {
	return &LruCache{
		lock:  &sync.RWMutex{},
		max:   max,
		count: 0,
		lru:   list.New(),
		data:  map[interface{}]interface{}{},
	}
}

func (c *LruCache) Put(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	var e *list.Element
	_e, ok := c.data[key]
	if ok {
		e = _e.(*list.Element)
		e.Value = []interface{}{key, value}
		c.lru.MoveToFront(e)
	} else {
		if c.count >= c.max {
			e = c.lru.Back()
			delete(c.data, e.Value.([]interface{})[0])
			c.lru.Remove(e)
			c.count--
		}

		e = c.lru.PushFront([]interface{}{key, value})
		c.data[key] = e
		c.count++
	}
}

func (c *LruCache) Get(key interface{}) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_e, ok := c.data[key]
	if !ok {
		return nil
	} else {
		e := _e.(*list.Element)
		c.lru.MoveToFront(e)
		return e.Value.([]interface{})[1]
	}
}

func (c *LruCache) Exists(key interface{}) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.data[key]
	return ok
}

func (c *LruCache) Del(key interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	var e *list.Element
	_e, ok := c.data[key]
	if !ok {
		return false
	} else {
		e = _e.(*list.Element)
		delete(c.data, e.Value.([]interface{})[0])
		c.lru.Remove(e)
		c.count--
		return true
	}
}

// EOF
