package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	modify(s)

	fmt.Println(s)          // ???
	fmt.Println(s[:cap(s)]) // ???
}

func modify(s []int) {
	s = append(s, 4)
	s[0] = 10
	s = append(s, 5)
}
