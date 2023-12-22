package cache

import "log"

type Cache interface {
	Set(key string, value interface{})
	Delete(key string)
	Get(key string) interface{}
}
type MemoryCache struct {
	data map[string]interface{}
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]interface{}),
	}
}

func (m *MemoryCache) Set(key string, value interface{}) {
	m.data[key] = value
}
func (m *MemoryCache) Delete(key string) {
	_, ok := m.data[key]
	if !ok {
		log.Fatal("not found cache")
	}
	delete(m.data, key)
}
func (m *MemoryCache) Get(key string) interface{} {
	value, ok := m.data[key]
	if !ok {
		log.Fatal("not found cache")
		return nil
	}
	return value
}
