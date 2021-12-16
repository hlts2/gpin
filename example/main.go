package main

import (
	"fmt"
	"time"

	"github.com/hlts2/gpin"
)

type Counter struct {
	l gpin.Spinlock
	v map[string]int
}

func (c *Counter) Increment(key string) {
	c.l.Lock()
	c.v[key]++
	c.l.Unlock()
}

func (c *Counter) Value(key string) int {
	c.l.Lock()
	defer c.l.Unlock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Increment("example")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("example")) // 100
}
