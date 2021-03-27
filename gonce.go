package gonce

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	mutex sync.Mutex
	done  uint32
}

func (once *Once) Do(callable func()) {
	if atomic.LoadUint32(&once.done) == 1 {
		return
	}

	once.mutex.Lock()
	defer once.mutex.Unlock()
	if once.done == 0 {
		defer atomic.StoreUint32(&once.done, 1)
		callable()
	}
}
