package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numbers := make([]int, 0, 10)
	var mux sync.Mutex

	for i := 0; i < 10; i++ {
		go func(v int) {
			mux.Lock()
			numbers = append(numbers, v)
			mux.Unlock()
		}(i)

	}
	time.Sleep(time.Second)
	fmt.Println(numbers)
}
