package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", p)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 350)
	fmt.Printf("%x\n", "f")
	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	// For basic string printing use %s.
	fmt.Printf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q.
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
}
