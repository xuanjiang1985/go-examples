package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test5.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
}
