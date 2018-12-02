package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCount(t *testing.T) {
	scenarios := []struct {
		input string
		out   map[int][]rune
	}{
		{input: "bababc", out: map[int][]rune{
			1: []rune("c"),
			3: []rune("b"),
			2: []rune("a"),
		}},
		{input: "abcccd", out: map[int][]rune{
			3: []rune("c"),
			1: []rune("abd"),
		}},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, count(scenario.input), scenario.out, scenario.input)
	}
}

func TestChecksum(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	assert.Equal(t, checksum(input), 12)
}
