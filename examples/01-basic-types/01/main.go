package main

import (
	"fmt"
)

func main() {
	var x1 int
	x2 := 0

	fmt.Println(x1)
	fmt.Println(x2)

	var y1 *int
	y2 := &x1
	y3 := new(int)

	x1 = 1

	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
}
