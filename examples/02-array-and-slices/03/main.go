package main

import (
	"fmt"
)

type BigStruct struct {
	Array [1000]int64
}

type SmallStruct struct {
	Array [10]int64
}

// go run -gcflags "-m" main.go
func main() {
	var a [10000]BigStruct
	foo(a)
	fmt.Println(a[0].Array[0])

	var b [10]SmallStruct
	bar(b)
	fmt.Println(b[0].Array[0])
}

//go:noinline
func foo(a [10000]BigStruct) {
	fmt.Println("test foo")
	a[0].Array[0] = 1
}

//go:noinline
func bar(a [10]SmallStruct) {
	fmt.Println("test bar")
	a[0].Array[0] = 1
}
