package main

import TheType "github.com/yuanqijing/sampleCode/basic/cmd/interface/type"

type D = TheType.D

const (
	d1 D = iota
)

type Arg = TheType.Arg

func main() {
	t := TheType.Mytype{}
	t.Hello(Arg{1, 2, d1})
}
