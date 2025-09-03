package pokecache

import(
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}

type Cache struct {
	data	map[string]cacheEntry
	mu 		*sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c :=  Cache {
		data:	make(map[string]cacheEntry),
		mu:   	&sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c

}

func (c *Cache)Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry {
		val:		val,
		createdAt: 	time.Now().UTC(),

	}
}

func (c *Cache)Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val.val, ok
}

func (c *Cache)reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}