package main

import (
	"fmt"
	"runtime"
	"time"
)

var i = 0

func main() {
	runtime.GOMAXPROCS(2)
	go func() {
		for {
			fmt.Println("i am here", i)
			time.Sleep(time.Second)
		}
	}()
	for {
		i += 1
		time.Sleep(time.Nanosecond)
	}
}
