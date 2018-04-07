package main

import (
	"fmt"
	"time"
)

func main(){
	// 超时机制
	timeout := make(chan bool)
	ch := make(chan int)
	go func(){
		time.Sleep(time.Second)
		timeout <- true
	}()
	go func(){
		ch <- 1
	}()
	for {
		select {
		case <-ch:
			fmt.Println("acceptted from ch")
			return
		case <-timeout:
			fmt.Println("time out, progrem exit")
			return
		}

	}
}