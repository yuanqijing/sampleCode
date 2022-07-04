package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 3)
		close(c)
	}()

	a := c
	select {
	case <-time.After(time.Second * 4):
		fmt.Println("yes")
	case _, ok := <-a:
		fmt.Println(ok)
		fmt.Println("not")
	}
	b, ok := <-a
	fmt.Println(ok)
	fmt.Println(b)
	b = <-a
	fmt.Println(b)
	b = <-a
	fmt.Println(b)
}
