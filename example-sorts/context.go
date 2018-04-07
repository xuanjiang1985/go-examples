package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("waited for 1 sec")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}