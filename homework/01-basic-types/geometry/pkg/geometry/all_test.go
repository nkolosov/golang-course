package geometry_test

import (
	"geometry-app/pkg/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Parallel()

	// Test Point
	t.Run("Point", func(t *testing.T) {
		t.Parallel()

		p := geometry.Point{X: 1, Y: 2}
		assert.InEpsilon(t, 1.0, p.X, 1e-9)
		assert.InEpsilon(t, 2.0, p.Y, 1e-9)
	})

	// Test Distance
	t.Run("Distance", func(t *testing.T) {
		t.Parallel()

		p1 := geometry.Point{X: 0, Y: 0}
		p2 := geometry.Point{X: 3, Y: 4}
		distance := p1.DistanceTo(p2)
		assert.InEpsilon(t, 5.0, distance, 1e-9)
	})

	// Test Polygon
	t.Run("Polygon", func(t *testing.T) {
		t.Parallel()

		p := geometry.Polygon{
			Points: []geometry.Point{
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
		t.Parallel()

		c := geometry.Circle{
			Center: geometry.Point{X: 1, Y: 1},
			Radius: 2,
		}
		assert.Equal(t, geometry.Point{X: 1, Y: 1}, c.Center)
		assert.InEpsilon(t, 2.0, c.Radius, 1e-9)
	})

	// Test interfaces
	t.Run("Interfaces", func(t *testing.T) {
		t.Parallel()

		var _ geometry.Figure = geometry.Circle{}
		var _ geometry.Figure = geometry.Polygon{}
	})
}
