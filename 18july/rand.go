package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go’s math/rand package provides pseudorandom number generation.

// Random function to illustrate pseudorandom number generation.
func main() {
	// fmt.Print(rand.Intn(100), ",")
	// fmt.Print(rand.Intn(100))

	// fmt.Println()

	// // rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	// fmt.Println(rand.Float64())

	// s2 := rand.NewSource(42)
	// r2 := rand.New(s2)
	// fmt.Print(r2.Intn(100), ",")
	// fmt.Print(r2.Intn(100))

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Call the resulting rand.Rand just like the functions on the rand package.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
}
