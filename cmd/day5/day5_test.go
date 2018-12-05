package main

import (
	"github.com/stretchr/testify/assert"
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

func TestRemoveCaseInsensitive(t *testing.T) {
	polymer := "dabAcCaCBAcCcaDA"
	scenarios := []struct {
		in  string
		out string
	}{
		{in: "a", out: "dbcCCBcCcD"},
		{in: "A", out: "dbcCCBcCcD"},
		{in: "b", out: "daAcCaCAcCcaDA"},
		{in: "B", out: "daAcCaCAcCcaDA"},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, removeCaseInsensitive(polymer, scenario.in), scenario.out, scenario.out)
	}
}

func TestShortestWithRemoval(t *testing.T) {
	in := "dabAcCaCBAcCcaDA"
	assert.Equal(t, len(ShortestWithRemoval(in)), 4)
}
