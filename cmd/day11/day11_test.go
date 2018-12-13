package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPower(t *testing.T) {
	scenarios := []struct {
		x int
		y int
		serial int
		power int
	} {
		{ 3, 5, 8, 4},
		{ 122, 79, 57, -5},
		{ 217, 196, 39, 0},
		{ 101, 153, 71, 4},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.power, power(scenario.x, scenario.y, scenario.serial))
	}
}