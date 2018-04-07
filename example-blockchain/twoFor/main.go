package main

import (
	"fmt"
	"time"
)

type demo struct {
	Value string
}

var pt *demo

func main() {
	//spinner(100 * time.Millisecond)
	pt = &demo{Value: "val"}
	fmt.Println(pt)
}

func spinner(delay time.Duration) {
	str := `-\|/`
	for {
		for _, r := range str {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
