package lruttl

import (
	"container/list"
	"sync"
	"time"
)

type item struct {
	key       string
	value     interface{}
	refreshed time.Time
	ttl       time.Duration
}

type LRU struct {
	capacity int
	ttl      time.Duration
	getFunc  func(string) (interface{}, time.Duration, error)
	items    map[string]*list.Element
	queue    *list.List
	mu       sync.RWMutex
}

func NewCache(capacity int, globalTTL time.Duration, getFunc func(string) (interface{}, time.Duration, error)) *LRU {
	return &LRU{
		capacity: capacity,
		ttl:      globalTTL,
		getFunc:  getFunc,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRU) set(key string, value interface{}, ttl time.Duration) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*item).value = value
		element.Value.(*item).refreshed = time.Now()
		element.Value.(*item).ttl = ttl
		return true
	}

	if c.queue.Len() >= c.capacity {
		c.purge()
	}

	i := &item{
		key:       key,
		value:     value,
		refreshed: time.Now(),
		ttl:       ttl,
	}

	element := c.queue.PushFront(i)
	c.items[i.key] = element

	return false
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*item)
		delete(c.items, item.key)
	}
}

func (c *LRU) Get(key string) (interface{}, error) {
	c.mu.RLock()
	element, exists := c.items[key]
	c.mu.RUnlock()
	if exists == false || (c.ttl != 0 && time.Now().After(element.Value.(*item).refreshed.Add(c.ttl))) || (element.Value.(*item).ttl != 0 && time.Now().After(element.Value.(*item).refreshed.Add(element.Value.(*item).ttl))) {
		i, ttl, err := c.getFunc(key)
		if err != nil {
			return i, err
		}
		c.set(key, i, ttl)
		return i, nil
	}

	c.queue.MoveToFront(element)
	return element.Value.(*item).value, nil
}
