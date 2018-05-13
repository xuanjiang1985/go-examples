package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var (
		total int32 = 0
	)
	for i := 0; i < 10; i++ {
		go func() {
			for {

				atomic.AddInt32(&total, 1)
				fmt.Println("now total is: ", total)
				time.Sleep(time.Millisecond)

			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("finally total is: ", atomic.LoadInt32(&total))
}
