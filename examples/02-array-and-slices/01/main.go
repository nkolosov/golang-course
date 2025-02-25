package main

import (
	"fmt"
)

func main() {
	var s []*int

	for _, v := range []int{1, 2, 3, 4} {
		s = append(s, &v)
	}

	for _, v := range s {
		fmt.Printf("value is %d\n", *v)
	}
}
