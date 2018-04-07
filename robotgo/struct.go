package main

import (
	"fmt"
)

type person struct {
	string
	int
}

func main() {
	a := &struct {
		Name string
		Age  int
	}{
		Name: "zhougang",
		Age:  19,
	}
	fmt.Println(a)
}
