package gpin

import (
	"sync"
	"sync/atomic"
	"testing"
)

const parallelism = 300

func BenchmarkMutex(b *testing.B) {
	var want int64
	got := 0
	mu := sync.Mutex{}

	b.SetParallelism(parallelism)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&want, 1)
			mu.Lock()
			got++
			mu.Unlock()
		}
	})
	if int(want) != got {
		b.Errorf("want %d, but got: %d", want, got)
	}
}

func BenchmarkSpinlock(b *testing.B) {
	var want int64
	got := 0
	l := &Spinlock{}

	b.SetParallelism(parallelism)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&want, 1)
			l.Lock()
			got++
			l.Unlock()
		}
	})
	if int(want) != got {
		b.Errorf("want %d, but got: %d", want, got)
	}
}
