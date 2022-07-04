package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type mystruct struct {
	Name  string
	slots []int64
	conn  *mystruct
	Map   map[int]string
	Map2  map[int]string
	Map3  map[int]string
}

func main() {
	m := &mystruct{
		Name: "helo",
		Map:  make(map[int]string),
	}
	m.Map[98] = "jj"
	str := spew.Sdump(m)
	fmt.Println(str)
}
