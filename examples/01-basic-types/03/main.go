package main

import (
	"fmt"
)

type A struct {
}

func (a *A) Foo() {
	fmt.Printf("A\n")
}

type B struct {
}

func (b *B) Foo() {
	fmt.Printf("B\n")
}

type C struct {
	A
}

func main() {
	c := &C{}
	c.Foo()
}
