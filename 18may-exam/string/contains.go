package main

import (
	"fmt"
	"strings"
)

func main() {
	p := "zhou gang is 33"
	//c := "zhou"
	b := strings.ContainsAny(p, "u o")
	fmt.Printf("%t\n", b)
	fmt.Println(strings.Count("five", ""))
}
