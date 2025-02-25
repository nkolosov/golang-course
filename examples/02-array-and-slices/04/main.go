package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println(a[4])

	s := []interface{}{"test", a, false, 123}
	fmt.Println(s[3])
}
