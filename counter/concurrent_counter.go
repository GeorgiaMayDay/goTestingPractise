package counter

type Counter struct {
	count int
}

func (c *Counter) Inc() {
	c.count++
}

func (c *Counter) Current() (currentNum int) {
	return c.count
}
