package gpin

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// Spinlock provides the spinlock.
type Spinlock struct {
	lock uint32
	_    sync.Mutex
}

// NewSpinlock returns sync.Locker implementation.
func NewSpinlock() sync.Locker {
	return &Spinlock{}
}

// Lock locks s.
// If the lock is already in use, the calling goroutine
// blocks until unlock.
func (s *Spinlock) Lock() {
	for {
		if atomic.CompareAndSwapUint32(&s.lock, 0, 1) {
			return
		}
		runtime.Gosched()
	}
}

// Unlock unlocks s.
// It is allowed for one goroutine to lock and then
// arrange for another goroutine to unlock it.
func (s *Spinlock) Unlock() {
	atomic.StoreUint32(&s.lock, 0)
}
