package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReact(t *testing.T) {
	scenarios := []struct {
		in  string
		out string
	}{
		{in: "aA", out: ""},
		{in: "abBA", out: "aA"},
		{in: "abAB", out: "abAB"},
		{in: "aabAAB", out: "aabAAB"},
		{in: "UDdHWwgx", out: "UHWwgx"},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.out, react(scenario.in), scenario.in)
	}
}

func TestReduce(t *testing.T) {
	scenarios := []struct {
		in  string
		out string
	}{
		{in: "aA", out: ""},
		{in: "abBA", out: ""},
		{in: "abAB", out: "abAB"},
		{in: "aabAAB", out: "aabAAB"},
		{in: "UDdHWwgx", out: "UHgx"},
		{in: "UDdHtNnxHpWAAafLUuLlDOIouKWWWjlWwgxQCK", out: "UHtxHpWAfLDOIouKWWWjlgxQCK"},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, reduce(scenario.in), scenario.out, scenario.in)
	}
}
