package geometry_test

import (
	"geometry-app/pkg/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointDistanceTo(t *testing.T) {
	t.Parallel()

	t.Run("distance between same points is zero", func(t *testing.T) {
		t.Parallel()

		p1 := geometry.Point{X: 0, Y: 0}
		p2 := geometry.Point{X: 0, Y: 0}
		distance := p1.DistanceTo(p2)
		assert.InEpsilon(t, 0.0, distance, 1e-9)
	})

	t.Run("distance between different points", func(t *testing.T) {
		t.Parallel()

		p1 := geometry.Point{X: 0, Y: 0}
		p2 := geometry.Point{X: 3, Y: 4}
		distance := p1.DistanceTo(p2)
		assert.InEpsilon(t, 5.0, distance, 1e-9)
	})

	t.Run("distance with negative coordinates", func(t *testing.T) {
		t.Parallel()

		p1 := geometry.Point{X: -1, Y: -1}
		p2 := geometry.Point{X: 2, Y: 3}
		distance := p1.DistanceTo(p2)
		assert.InEpsilon(t, 5.0, distance, 1e-9)
	})
}
