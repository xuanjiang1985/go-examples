package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(get())
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}
}

func get() (a int) {
	v := rand.Intn(1000)
	defer func() {
		if e := recover(); e != nil {
			a = v
		}
	}()
	panic(0)
}
