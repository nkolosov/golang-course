package geometry

import "math"

// Point represents a point in 2D space
type Point struct {
	X float64
	Y float64
}

// DistanceTo calculates the Euclidean distance between this point and another point
func (p Point) DistanceTo(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}
