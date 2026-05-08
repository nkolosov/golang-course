package geometry_test

import (
	"geometry-app/pkg/geometry"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleArea(t *testing.T) {
	t.Parallel()

	t.Run("circle area", func(t *testing.T) {
		t.Parallel()
		c := geometry.Circle{
			Center: geometry.Point{X: 0, Y: 0},
			Radius: 2,
		}
		area := c.Area()
		assert.InEpsilon(t, 4*math.Pi, area, 1e-9)
	})

	t.Run("circle with zero radius", func(t *testing.T) {
		t.Parallel()
		c := geometry.Circle{
			Center: geometry.Point{X: 0, Y: 0},
			Radius: 0,
		}
		area := c.Area()
		assert.InEpsilon(t, 0.0, area, 1e-9)
	})
}

func TestCirclePerimeter(t *testing.T) {
	t.Parallel()

	t.Run("circle perimeter", func(t *testing.T) {
		t.Parallel()
		c := geometry.Circle{
			Center: geometry.Point{X: 0, Y: 0},
			Radius: 2,
		}
		perimeter := c.Perimeter()
		assert.InEpsilon(t, 4*math.Pi, perimeter, 1e-9)
	})

	t.Run("circle with zero radius", func(t *testing.T) {
		t.Parallel()
		c := geometry.Circle{
			Center: geometry.Point{X: 0, Y: 0},
			Radius: 0,
		}
		perimeter := c.Perimeter()
		assert.InEpsilon(t, 0.0, perimeter, 1e-9)
	})
}

func TestCircleContains(t *testing.T) {
	t.Parallel()

	t.Run("point inside circle", func(t *testing.T) {
		t.Parallel()
		c := geometry.Circle{
			Center: geometry.Point{X: 0, Y: 0},
			Radius: 5,
		}
		point := geometry.Point{X: 3, Y: 4}
		contains := c.Contains(point)
		assert.True(t, contains)
	})

	t.Run("point outside circle", func(t *testing.T) {
		t.Parallel()
		c := geometry.Circle{
			Center: geometry.Point{X: 0, Y: 0},
			Radius: 5,
		}
		point := geometry.Point{X: 6, Y: 0}
		contains := c.Contains(point)
		assert.False(t, contains)
	})

	t.Run("point on circle boundary", func(t *testing.T) {
		t.Parallel()
		c := geometry.Circle{
			Center: geometry.Point{X: 0, Y: 0},
			Radius: 5,
		}
		point := geometry.Point{X: 5, Y: 0}
		contains := c.Contains(point)
		assert.True(t, contains)
	})
}
