package TheType

import "fmt"

type D int

type Arg struct {
	A int
	B int
	C D
}

type Myinterface interface {
	Hello(arg Arg)
}

type Mytype struct {
	a int
}

func (m *Mytype) Hello(arg Arg) {
	fmt.Printf("hello")
}
