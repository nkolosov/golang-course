package geometry

// Figure interface defines methods for geometric figures
type Figure interface {
	Area() float64
	Perimeter() float64
	Contains(point Point) bool
}

const epsilon = 1e-9
