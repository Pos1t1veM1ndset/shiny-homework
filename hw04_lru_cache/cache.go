package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mux      *sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, v interface{}) (ok bool) {
	//mux := &sync.RWMutex{}
	c.mux.Lock()
	val, ok := c.items[key]
	if !ok {
		if c.queue.Len() == c.capacity {
			lastInQueue := c.queue.Back()
			c.queue.Remove(lastInQueue)
			delete(c.items, (lastInQueue.Value).(cacheItem).key)
		}
		tmp := c.queue.PushFront(cacheItem{key, v})
		c.items[key] = tmp
		c.mux.Unlock()
		return
	}
	c.items[key].Value = cacheItem{key, v}
	c.queue.MoveToFront(val)
	c.mux.Unlock()
	return
}

func (c lruCache) Get(key Key) (value interface{}, ok bool) {
	//mux := &sync.RWMutex{}
	c.mux.Lock()
	val, ok := c.items[key]
	if ok {
		value = (val.Value).(cacheItem).value
		c.queue.MoveToFront(val)
		c.mux.Unlock()
		return
	}
	value = nil
	c.mux.Unlock()
	return
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mux:      &sync.Mutex{},
	}
}
