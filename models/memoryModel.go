package models

import "sync"

type MemoryCache struct {
	Cache map[string]interface{}
	Mutex sync.RWMutex
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	data, exists := c.Cache[key]
	return data, exists
}

func (c *MemoryCache) Set(key string, data interface{}) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Cache[key] = data
}
