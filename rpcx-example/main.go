package main

import (
	"fmt"
)

func main() {

	// s := []byte("")
	// s1 := append(s, 'a', 'c', 'd')

	// s2 := append(s, 'b')
	// fmt.Println('a')
	// fmt.Println(string(s1))
	// fmt.Println(string(s2))
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	res := 0
	close(ch)
	for v := range ch {
		res += v
	}
	fmt.Println(res)

}
