package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxa, cancela := context.WithCancel(ctx)
	go a(ctxa)
	ctxb, cancelb := context.WithCancel(ctxa)
	go b(ctxb)

	// 先 cancel a 再 cancel b
	cancela()
	time.Sleep(time.Second * 3)
	fmt.Printf("=======")
	cancelb()
	/*
		b canceled
		a canceled
		=======
	*/

	// 先 cancel b 再 cancel a
	cancelb()
	time.Sleep(time.Second * 3)
	fmt.Printf("=======")
	cancela()
	time.Sleep(time.Second * 1)
	/*
		b canceled
		=======
		a canceled
	*/
}

func a(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("a canceled")
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

func b(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("b canceled")
			return
		default:
			time.Sleep(time.Second)
		}
	}
}
