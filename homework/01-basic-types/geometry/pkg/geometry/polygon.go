package geometry

import (
	"math"
)

const MinPoints = 3

// Polygon represents a polygon defined by a slice of points
type Polygon struct {
	Points []Point
}

// Area calculates the area of the polygon using the shoelace formula
func (p Polygon) Area() float64 {
	if len(p.Points) < MinPoints {
		return 0
	}

	var area float64
	n := len(p.Points)

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += p.Points[i].X * p.Points[j].Y
		area -= p.Points[j].X * p.Points[i].Y
	}

	return math.Abs(area) / 2
}

// Perimeter calculates the perimeter of the polygon
func (p Polygon) Perimeter() float64 {
	if len(p.Points) < MinPoints {
		return 0
	}

	var perimeter float64
	n := len(p.Points)

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		perimeter += p.Points[i].DistanceTo(p.Points[j])
	}

	return perimeter
}

// Contains checks if a point is inside the polygon using ray casting algorithm
func (p Polygon) Contains(point Point) bool {
	if len(p.Points) < MinPoints {
		return false
	}

	intersections := 0
	n := len(p.Points)

	for i := 0; i < n; i++ {
		j := (i + 1) % n

		// Check if point is on the edge
		if isPointOnLineSegment(p.Points[i], p.Points[j], point) {
			return true
		}

		// Ray casting algorithm
		if ((p.Points[i].Y > point.Y) != (p.Points[j].Y > point.Y)) &&
			(point.X < (p.Points[j].X-p.Points[i].X)*(point.Y-p.Points[i].Y)/(p.Points[j].Y-p.Points[i].Y)+p.Points[i].X) {
			intersections++
		}
	}

	return intersections%2 == 1
}

// Helper function to check if point is on line segment
func isPointOnLineSegment(a, b, c Point) bool {
	// Check if point c lies on the line segment from a to b
	// Using cross product to check collinearity
	crossProduct := (c.Y-a.Y)*(b.X-a.X) - (c.X-a.X)*(b.Y-a.Y)
	if math.Abs(crossProduct) > epsilon {
		return false
	}

	// Check if point c is within the bounding box of segment ab
	minX := math.Min(a.X, b.X)
	maxX := math.Max(a.X, b.X)
	minY := math.Min(a.Y, b.Y)
	maxY := math.Max(a.Y, b.Y)

	return c.X >= minX && c.X <= maxX && c.Y >= minY && c.Y <= maxY
}
