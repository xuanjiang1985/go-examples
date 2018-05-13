package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./main.go")
	defer f.Close()

	buff := make([]byte, 1024)
	f.Read(buff)
	fmt.Println(string(buff))
}
