package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error message 500")
	fmt.Printf("%s\n", err)
}
