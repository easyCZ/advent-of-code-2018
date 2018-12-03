package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCountOverlap(t *testing.T) {
	grid := &Grid{
		claims: []*Claim{
			{ id: "1", left: 1, top: 3, width: 4, height: 4},
			{ id: "2", left: 3, top: 1, width: 4, height: 4},
			{ id: "3", left: 5, top: 5, width: 2, height: 2},
		},
	}

	assert.Equal(t, grid.countOverlap(2), 4)
}
