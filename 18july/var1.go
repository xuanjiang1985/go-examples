package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var i = 0

func main() {
	runtime.GOMAXPROCS(2)
	mtx := sync.RWMutex{}
	go func() {
		for {
			mtx.RLock()
			fmt.Println("i am here", i)
			mtx.RUnlock()
			time.Sleep(time.Second)
		}
	}()
	for {
		mtx.Lock()
		i += 1
		mtx.Unlock()
		time.Sleep(time.Nanosecond)
	}
}
