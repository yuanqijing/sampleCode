package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	Process(ctx)
}

func Process(ctx context.Context) {
	finish := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		finish <- 1
	}()
	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
	case <-finish:
		fmt.Println("finished")
	}
}
