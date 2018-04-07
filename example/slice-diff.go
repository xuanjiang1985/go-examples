package main

import "fmt"

func main() {
	var a []byte
	fmt.Println(a, len(a), cap(a))

	b := make([]byte, 0)
	fmt.Println(b, len(b), cap(b))
}
