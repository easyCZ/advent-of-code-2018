package main

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGrid_Influence(t *testing.T) {
	g := &Grid{
		points: []*Point{
			{1, 1},
			{1, 6},
			{8, 3},
			{3, 4},
			{5, 5},
			{8, 9},
		},
	}

	plot := g.Plot()

	//assert.Equal(t, g.influence(g.points[0], plot), 15)
	assert.Equal(t, g.influence(g.points[4], plot), 17)
}

func TestMaxInfluence(t *testing.T) {
	g := &Grid{
		points: []*Point{
			{1, 1},
			{1, 6},
			{8, 3},
			{3, 4},
			{5, 5},
			{8, 9},
		},
	}

	plot := g.Plot()
	for _, p := range plot {
		fmt.Println(p)
	}
	maxPoint, maxInfluence := g.MaxInfluence()
	assert.Equal(t, g.points[4], maxPoint)
	assert.Equal(t, maxInfluence, 17)
}