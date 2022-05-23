package main

import (
	"fmt"
)

type I int

func (i I) String() string {
	return "I"
}

type S string

func (s S) String() string {
	return "S"
}

type F float32

func (f F) String() string {
	return "F"
}

type Stringer interface {
	String() string
}

func main() {

	var i Stringer = I(100)
	f(i)

	var s Stringer = S("test")
	f(s)

	var fl Stringer = F(0.1)
	f(fl)
}

func f(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println("int: ", int(v))
	case S:
		fmt.Println("string: ", string(v))
	case F:
		fmt.Println("float: ", float32(v))
	}
}
