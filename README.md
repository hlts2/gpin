# gpin
A [spinlock](https://en.wikipedia.org/wiki/Spinlock) implementation for Go. 

## Requirement

Go 1.17


## Installation
```shell
go get github.com/hlts2/gpin
```

## Example

```go
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

func (c *Counter) Get(key string) int {
	c.l.Lock()
	defer c.l.Unlock()
	return c.v[key]
}

func main() {
	c := Counter{
		v: make(map[string]int),
	}
	for i := 0; i < 100; i++ {
		go c.Increment("example")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Get("example")) // 100
}

```

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/hlts2/gpin
cpu: 11th Gen Intel(R) Core(TM) i9-11900K @ 3.50GHz
BenchmarkMutex
BenchmarkMutex-16       	10193804	       135.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSpinlock
BenchmarkSpinlock-16    	43065172	        30.98 ns/op	       0 B/op	       0 allocs/op
```
