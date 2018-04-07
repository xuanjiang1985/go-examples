// package main

// import (
// 	"math/rand"
// 	"sync"
// 	"fmt"
// )

// const N = 10

// func main() {
// 	 m := make(map[int]int)

// 	 // 10个队列来阻塞，直到全部完成
// 	 wg := &sync.WaitGroup{}
// 	 wg.Add(N)
	 
// 	 mu := &sync.Mutex{}
// 	 for i := 0; i < N; i++ {
// 	 	go func() {
// 			 defer wg.Done()
// 			 mu.Lock()
// 			 m[rand.Int()] = rand.Int()
// 			 mu.Unlock()
// 	 	}()
// 	 }
// 	 wg.Wait()
// 	fmt.Println(m)
// }