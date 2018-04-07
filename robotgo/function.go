package main

import (
	"fmt"
)

func main() {
	a()
	b()
	c()
}

func a() {
	fmt.Println("Func a")
}

func b() {
	fmt.Println("Func b")
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("Recover in b")
		}
	}()
	panic("Panic in b")
}

func c() {
	fmt.Println("Func c")
}
