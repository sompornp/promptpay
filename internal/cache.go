package internal

import (
	"sync"
)

// Cache structure
type Cache struct {
	lock  sync.RWMutex
	store map[string]interface{}
}

// NewCache creates a new Cache instance
func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

// Set a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.store[key] = value
}

// Get a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	val, ok := c.store[key]
	return val, ok
}

// Delete a value from the cache
func (c *Cache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.store, key)
}
