package main

import (
	"fmt"
)

func main() {
	fmt.Println(echo())
}

func echo() string {
	name := "zhougang"
	defer fmt.Println("hi 1")
	defer fmt.Println("hi 2")
	return name
}