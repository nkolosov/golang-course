package main

import "fmt"

func main() {
	// Slice of structs
	type Point struct {
		X, Y int
	}

	points := []Point{{1, 2}, {3, 4}, {5, 6}}

	// Ситуация 1
	for _, p := range points {
		p.X += 10
	}
	fmt.Println(points) // ???

	// Ситуация 2
	for i := range points {
		points[i].Y += 10
	}
	fmt.Println(points) // ???

	// Ситуация 3: map of slices
	m := make(map[string][]int)
	m["key"] = []int{1, 2, 3}

	modifyMap(m)
	fmt.Println(m["key"]) // ???

	// Ситуация 4: append в map
	m2 := map[string][]int{"a": {1, 2}}
	val := m2["a"]
	val = append(val, 3)
	fmt.Println(m2["a"]) // ???
}

func modifyMap(m map[string][]int) {
	m["key"][0] = 100
	m["key"] = append(m["key"], 4)
}
