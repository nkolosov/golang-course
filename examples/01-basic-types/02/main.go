package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	A int8
	B int16
	C int64
}

type B struct {
	A int8
	С int64
	B int16
}

func main() {
	a := A{}
	b := B{}

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
