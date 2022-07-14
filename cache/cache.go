package cache

import "errors"

type Cache struct {
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c *Cache) Find(key string) (interface{}, error) {
	result, ok := c.cache[key]
	if !ok {
		return nil, errors.New("Value not found")
	}
	return result, nil
}

func (c *Cache) Add(key string, value interface{}) {
	c.cache[key] = value
}

func (c *Cache) Del(key string) error {
	if _, ok := c.cache[key]; !ok {
		return errors.New("Key inexistent")
	}
	delete(c.cache, key)
	return nil
}
