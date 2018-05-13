package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./atomic.go")
	fmt.Println(err)
	fmt.Println(string(data))
}
