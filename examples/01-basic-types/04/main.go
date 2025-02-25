package main

import (
	"fmt"
	"unsafe"
)

type BigStruct struct {
	A, B, C, D int64
}

func main() {
	a := 1
	fmt.Println(unsafe.Sizeof(a)) // ?? bytes
	printSize(a)                  // ?? bytes
	fmt.Println()

	b := false
	fmt.Println(unsafe.Sizeof(b)) // ?? bytes
	printSize(b)                  // ?? bytes
	fmt.Println()

	c := BigStruct{}
	fmt.Println(unsafe.Sizeof(c)) // ?? bytes
	printSize(c)                  // ?? bytes
	fmt.Println()

	d := struct{}{}
	fmt.Println(unsafe.Sizeof(d)) // ?? bytes
	printSize(d)                  // ?? bytes
	fmt.Println()
}

func printSize(a interface{}) {
	fmt.Println(unsafe.Sizeof(a))
}
