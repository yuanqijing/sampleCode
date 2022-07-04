package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strconv.Itoa(67))
	str := "123 $eee  fff"
	rec := strings.ReplaceAll(str, "$eee", "ppp")
	fmt.Println(str)
	fmt.Println(rec)

	a := "sdf"
	b := "sdf"
	fmt.Println(a == b)
}
