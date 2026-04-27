package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	// Test Point
	t.Run("Point", func(t *testing.T) {
		p := Point{X: 1, Y: 2}
		assert.Equal(t, 1.0, p.X)
		assert.Equal(t, 2.0, p.Y)
	})

	// Test Distance
	t.Run("Distance", func(t *testing.T) {
		p1 := Point{X: 0, Y: 0}
		p2 := Point{X: 3, Y: 4}
		distance := p1.DistanceTo(p2)
		assert.Equal(t, 5.0, distance)
	})

	// Test Polygon
	t.Run("Polygon", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 2, Y: 0},
				{X: 2, Y: 2},
				{X: 0, Y: 2},
			},
		}
		assert.Len(t, p.Points, 4)
	})

	// Test Circle
	t.Run("Circle", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 1, Y: 1},
			Radius: 2,
		}
		assert.Equal(t, Point{X: 1, Y: 1}, c.Center)
		assert.Equal(t, 2.0, c.Radius)
	})

	// Test interfaces
	t.Run("Interfaces", func(t *testing.T) {
		var _ Figure = Circle{}
		var _ Figure = Polygon{}
	})
}
