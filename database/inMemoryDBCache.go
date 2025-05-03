package database

import (
	"runtime"
	"sync"
	"time"
)

const (
	DefaultDuration      time.Duration = 0
	DefaultCacheInterval time.Duration = 2 * time.Minute
)

var (
	once         sync.Once
	defaultCache *Cache
)

func GetCache() *Cache {
	items := make(map[string]Item)
	once.Do(func() {
		defaultCache = &Cache{
			items: items,
		}
		runJanitor(DefaultCacheInterval, defaultCache)
		runtime.SetFinalizer(defaultCache, stopExceution)
	})
	return defaultCache
}

type Item struct {
	value    interface{}
	duration int64
}

func (this Item) Expired() bool {
	if this.duration == 0 {
		return false
	}
	return time.Now().UnixNano() > this.duration
}

type Cache struct {
	items   map[string]Item
	mu      sync.RWMutex
	janitor *Janitor
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	var d int64
	if ttl == 0 {
		ttl = DefaultDuration
	}
	if ttl > 0 {
		d = time.Now().Add(ttl).UnixNano()
	}
	c.mu.RLock()
	c.items[key] = Item{
		value:    value,
		duration: d,
	}
	return
}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	if item.Expired() {
		c.mu.RUnlock()
		return nil, false
	}
	c.mu.RUnlock()
	return item.value, true
}

func (c *Cache) Delete(key string) bool {
	c.mu.RLock()
	if _, ok := c.items[key]; !ok {
		return false
	}
	delete(c.items, key)
	c.mu.RUnlock()
	return true
}

func (c *Cache) Len() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	return n
}

func (c *Cache) Flush() {
	c.mu.RLock()
	c.items = map[string]Item{}
	c.mu.RUnlock()
}

func (c *Cache) DeleteExpired() {
	c.mu.RLock()
	for k, v := range c.items {
		if v.Expired() {
			c.Delete(k)
		}
	}
	c.mu.RUnlock()
}

type Janitor struct {
	interval time.Duration
	stop     chan bool
}

func (j *Janitor) Run(c *Cache) {
	ticker := time.NewTicker(j.interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
		}
	}
}

func stopExceution(c *Cache) {
	c.janitor.stop <- true

}

// TODO we could design this better
func runJanitor(interval time.Duration, c *Cache) {
	janitor := &Janitor{
		interval: interval,
		stop:     make(chan bool),
	}
	c.janitor = janitor
	go janitor.Run(c)
}

// TODO set context for methods why
// NOTE janitor should be private
// TODO make it a Cache package
