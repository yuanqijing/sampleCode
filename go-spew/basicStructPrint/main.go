package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type mystruct struct {
	Name  string
	other []*mystruct
}

func main() {
	a := mystruct{Name: "1st layer"}
	b := mystruct{Name: "2nd layer", other: []*mystruct{&a}}
	c := mystruct{Name: "3rd layer", other: []*mystruct{&b}}
	d := mystruct{Name: "4th layer", other: []*mystruct{&c}}
	str := spew.Sdump(d)
	fmt.Println(str)
	//str = spew.Sprintf("%#+v", d)
	//fmt.Println(str)
}
