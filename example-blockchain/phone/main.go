package main

import (
	"fmt"
)

func main() {
	data := "acb"
	for _, v := range data {
		fmt.Println(string(v))
	}
}

func pass(ps *string) {
	*ps = "donal"
}
