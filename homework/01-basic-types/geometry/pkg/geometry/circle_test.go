package geometry

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleArea(t *testing.T) {
	t.Run("circle area", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 0, Y: 0},
			Radius: 2,
		}
		area := c.Area()
		assert.Equal(t, 4*math.Pi, area)
	})

	t.Run("circle with zero radius", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 0, Y: 0},
			Radius: 0,
		}
		area := c.Area()
		assert.Equal(t, 0.0, area)
	})
}

func TestCirclePerimeter(t *testing.T) {
	t.Run("circle perimeter", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 0, Y: 0},
			Radius: 2,
		}
		perimeter := c.Perimeter()
		assert.Equal(t, 4*math.Pi, perimeter)
	})

	t.Run("circle with zero radius", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 0, Y: 0},
			Radius: 0,
		}
		perimeter := c.Perimeter()
		assert.Equal(t, 0.0, perimeter)
	})
}

func TestCircleContains(t *testing.T) {
	t.Run("point inside circle", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 0, Y: 0},
			Radius: 5,
		}
		point := Point{X: 3, Y: 4}
		contains := c.Contains(point)
		assert.True(t, contains)
	})

	t.Run("point outside circle", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 0, Y: 0},
			Radius: 5,
		}
		point := Point{X: 6, Y: 0}
		contains := c.Contains(point)
		assert.False(t, contains)
	})

	t.Run("point on circle boundary", func(t *testing.T) {
		c := Circle{
			Center: Point{X: 0, Y: 0},
			Radius: 5,
		}
		point := Point{X: 5, Y: 0}
		contains := c.Contains(point)
		assert.True(t, contains)
	})
}
