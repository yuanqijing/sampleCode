package main

import "fmt"

func main() {
	va := "\"\""
	va = "'\"''''\"'"
	str := fmt.Sprintf("%s", va)
	fmt.Printf(str)
}
