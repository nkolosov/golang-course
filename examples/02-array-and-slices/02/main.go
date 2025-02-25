package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 3, 2, 1}
	ap(a)

	fmt.Println(a) // ???
}

func ap(a []int) {
	a[0] = 0
	a = append(a, 10)
}

/**
package main

import (
    "fmt"
)

func main() {
    fmt.Println(arr())
}

func arr() []int {
    a := []int{0, 1, 2, 3, 4}
    a[0] = 1
    a = append(a, 5)
    newArr := append(a, 6)
    a[0] = 10
    return newArr
}
*/
