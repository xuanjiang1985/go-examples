package main

import (
	"fmt"
)

func main() {
	mp1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(mp1)
	mp2 := map[string]int{}
	for k, v := range mp1 {
		mp2[v] = k
	}
	fmt.Println(mp2)
}
