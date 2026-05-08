package geometry_test

import (
	"geometry-app/pkg/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolygonArea(t *testing.T) {
	t.Parallel()

	t.Run("triangle area", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 4},
			},
		}
		area := p.Area()
		assert.InEpsilon(t, 6.0, area, 1e-9)
	})

	t.Run("square area", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 2, Y: 0},
				{X: 2, Y: 2},
				{X: 0, Y: 2},
			},
		}
		area := p.Area()
		assert.InEpsilon(t, 4.0, area, 1e-9)
	})

	t.Run("empty polygon", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{}
		area := p.Area()
		assert.InEpsilon(t, 0.0, area, 1e-9)
	})

	t.Run("polygon with less than 3 points", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			},
		}
		area := p.Area()
		assert.InEpsilon(t, 0.0, area, 1e-9)
	})
}

func TestPolygonPerimeter(t *testing.T) {
	t.Parallel()

	t.Run("square perimeter", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 2, Y: 0},
				{X: 2, Y: 2},
				{X: 0, Y: 2},
			},
		}
		perimeter := p.Perimeter()
		assert.InEpsilon(t, 8.0, perimeter, 1e-9)
	})

	t.Run("triangle perimeter", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 4},
			},
		}
		perimeter := p.Perimeter()
		assert.InEpsilon(t, 12.0, perimeter, 1e-9)
	})

	t.Run("empty polygon", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{}
		perimeter := p.Perimeter()
		assert.InEpsilon(t, 0.0, perimeter, 1e-9)
	})
}

func TestPolygonContains(t *testing.T) {
	t.Parallel()

	t.Run("point inside triangle", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 3},
			},
		}
		point := geometry.Point{X: 1, Y: 1}
		contains := p.Contains(point)
		assert.True(t, contains)
	})

	t.Run("point outside triangle", func(t *testing.T) {
		t.Parallel()
		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 3},
			},
		}
		point := geometry.Point{X: 2, Y: 2}
		contains := p.Contains(point)
		assert.False(t, contains)
	})

	t.Run("point on edge", func(t *testing.T) {
		t.Parallel()
		p := geometry.Polygon{
			Points: []geometry.Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 3},
			},
		}
		point := geometry.Point{X: 1, Y: 0} // On the edge
		contains := p.Contains(point)
		assert.True(t, contains)
	})

	t.Run("empty polygon", func(t *testing.T) {
		t.Parallel()
		p := geometry.Polygon{}
		point := geometry.Point{X: 1, Y: 1}
		contains := p.Contains(point)
		assert.False(t, contains)
	})
}
