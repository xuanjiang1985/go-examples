package main

import (
	"time"
	"fmt"
)

func main(){
	// 管道
	// num := make(chan int, 10)
	// go func(){
	// 	for x := 0; x < 10; x++ {
	// 		num<- x
	// 		time.Sleep(time.Second)
	// 	}
	// 	close(num)
	// }()

	// for y := range num {
	// 	fmt.Println(y)
	// }
	// 并发循环
	loopgo()
	time.Sleep(time.Second)
}

func loopgo(){
	for i := 0; i < 10; i++ {
		go func(a int){
			fmt.Println(a)
		}(i)
	}
}