package counter

import "sync"

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Current() (currentNum int) {
	return c.count
}
