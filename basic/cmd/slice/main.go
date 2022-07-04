package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("%v", slice[1:])
	fmt.Printf("%v", slice[len(slice)-1])

	slice2 := []int{1, 2, 3}
	fmt.Printf("%v", slice2[1:])
	fmt.Printf("%v", slice2[len(slice2)-1])
	slice2[1] = 4

}
