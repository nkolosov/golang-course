package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointDistanceTo(t *testing.T) {
	t.Run("distance between same points is zero", func(t *testing.T) {
		p1 := Point{X: 0, Y: 0}
		p2 := Point{X: 0, Y: 0}
		distance := p1.DistanceTo(p2)
		assert.Equal(t, 0.0, distance)
	})

	t.Run("distance between different points", func(t *testing.T) {
		p1 := Point{X: 0, Y: 0}
		p2 := Point{X: 3, Y: 4}
		distance := p1.DistanceTo(p2)
		assert.Equal(t, 5.0, distance)
	})

	t.Run("distance with negative coordinates", func(t *testing.T) {
		p1 := Point{X: -1, Y: -1}
		p2 := Point{X: 2, Y: 3}
		distance := p1.DistanceTo(p2)
		assert.Equal(t, 5.0, distance)
	})
}
