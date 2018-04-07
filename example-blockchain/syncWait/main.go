package main

import (
	"fmt"
	"sync"
)

func main() {
	//waitGroup
	// var wg sync.WaitGroup
	// for i := 0; i < 6; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(i)
	// 	wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	// mutex
	var sum int = 0
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10000; i++ {
			m.Lock()
			sum++
			m.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			m.Lock()
			sum++
			m.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(sum)
}
