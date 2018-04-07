package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	fmt.Println(strings.Contains(s, "1"))
	fmt.Println(strings.Index(s, "w"))
	ss := "1#2#345"
	spliteStr := strings.Split(ss, "#")
	fmt.Println(spliteStr)
}
