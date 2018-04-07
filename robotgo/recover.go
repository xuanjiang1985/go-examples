package main

import "fmt"

func recoverPanic() {
	func() {
		if r := recover(); r != nil {
			fmt.Println("recover value is", r)
		}
	}()

	panic("exception")
}

func main() {
	recoverPanic()
}
