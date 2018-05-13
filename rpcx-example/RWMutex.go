package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		state = map[string]int{
			"a": 1,
		}
		mux sync.Mutex

		muxOps uint64

		state2 = map[string]int{
			"a": 1,
		}
		rw    sync.RWMutex
		rwOps uint64
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				_ = state["a"]
				mux.Unlock()
				atomic.AddUint64(&muxOps, 1)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				rw.RLock()
				_ = state2["a"]
				rw.RUnlock()
				atomic.AddUint64(&rwOps, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(muxOps, rwOps)
}
