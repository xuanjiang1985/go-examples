package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("1")
	}()
	go func() {
		fmt.Println("2")
	}()
	go func() {
		fmt.Println("3")
	}()
	runtime.Gosched()
}
