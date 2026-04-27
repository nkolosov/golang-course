package main

import (
	"fmt"
)

func main() {
	var x1 int
	x2 := 0

	fmt.Println(x1) // noinline
	fmt.Println(x2) // noinline

	var y1 *int
	y2 := &x1
	y3 := new(int)

	x1 = 1

	fmt.Println(y1) // noinline
	fmt.Println(y2) // noinline
	fmt.Println(y3) // noinline
}
