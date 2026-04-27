package geometry

import "math"

// Circle represents a circle with center point and radius
type Circle struct {
	Center Point
	Radius float64
}

// Area calculates the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter (circumference) of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Contains checks if a point is inside the circle
func (c Circle) Contains(point Point) bool {
	distance := c.Center.DistanceTo(point)
	return distance <= c.Radius
}
