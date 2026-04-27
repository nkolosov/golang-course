package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolygonArea(t *testing.T) {
	t.Run("triangle area", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 4},
			},
		}
		area := p.Area()
		assert.Equal(t, 6.0, area)
	})

	t.Run("square area", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 2, Y: 0},
				{X: 2, Y: 2},
				{X: 0, Y: 2},
			},
		}
		area := p.Area()
		assert.Equal(t, 4.0, area)
	})

	t.Run("empty polygon", func(t *testing.T) {
		p := Polygon{}
		area := p.Area()
		assert.Equal(t, 0.0, area)
	})

	t.Run("polygon with less than 3 points", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			},
		}
		area := p.Area()
		assert.Equal(t, 0.0, area)
	})
}

func TestPolygonPerimeter(t *testing.T) {
	t.Run("square perimeter", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 2, Y: 0},
				{X: 2, Y: 2},
				{X: 0, Y: 2},
			},
		}
		perimeter := p.Perimeter()
		assert.Equal(t, 8.0, perimeter)
	})

	t.Run("triangle perimeter", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 4},
			},
		}
		perimeter := p.Perimeter()
		assert.Equal(t, 12.0, perimeter)
	})

	t.Run("empty polygon", func(t *testing.T) {
		p := Polygon{}
		perimeter := p.Perimeter()
		assert.Equal(t, 0.0, perimeter)
	})
}

func TestPolygonContains(t *testing.T) {
	t.Run("point inside triangle", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 3},
			},
		}
		point := Point{X: 1, Y: 1}
		contains := p.Contains(point)
		assert.True(t, contains)
	})

	t.Run("point outside triangle", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 3},
			},
		}
		point := Point{X: 2, Y: 2}
		contains := p.Contains(point)
		assert.False(t, contains)
	})

	t.Run("point on edge", func(t *testing.T) {
		p := Polygon{
			Points: []Point{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 0, Y: 3},
			},
		}
		point := Point{X: 1, Y: 0} // On the edge
		contains := p.Contains(point)
		assert.True(t, contains)
	})

	t.Run("empty polygon", func(t *testing.T) {
		p := Polygon{}
		point := Point{X: 1, Y: 1}
		contains := p.Contains(point)
		assert.False(t, contains)
	})
}
