package mysync

import (
	"sync/atomic"
)

type Counter struct {
	value atomic.Int64
}

func (c *Counter) Inc() {
	c.value.Add(1)
}

func (c *Counter) Value() int64 {
	return c.value.Load()
}

func NewCounter() *Counter {
	return &Counter{}
}
