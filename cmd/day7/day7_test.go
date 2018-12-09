package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	input := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`
	var b bytes.Buffer
	b.WriteString(input)

	g := parse(&b)

	assert.Equal(t, []string{"A", "F"}, g.data["C"])
	assert.Equal(t, []string{"B", "D"}, g.data["A"])
	assert.Equal(t, []string{"E"}, g.data["D"])
	assert.Equal(t, []string{"E"}, g.data["B"])
	assert.Equal(t, []string{"E"}, g.data["F"])
}

func TestOrder(t *testing.T) {
	g := NewGraph(map[string][]string{
		"C": {"A", "F"},
		"A": {"B", "D"},
		"D": {"E"},
		"B": {"E"},
		"F": {"E"},
		"E": {},
	})

	//assert.Equal(t, "C", g.Root())
	assert.Equal(t, []string{"C", "A", "B", "D", "F", "E"}, g.Order())
}

func TestExecutePlan(t *testing.T) {
	g := NewGraph(map[string][]string{
		"C": {"A", "F"},
		"A": {"B", "D"},
		"D": {"E"},
		"B": {"E"},
		"F": {"E"},
		"E": {},
	})

	//assert.Equal(t, "C", g.Root())
	baseStepCost = 0
	time := ExecutePlan(g, 2)
	assert.Equal(t, 15, time)
}


func TestExclude(t *testing.T) {
	assert.Equal(t, []string{"A", "B", "D"}, exclude([]string{"A", "B", "C", "D"}, "C"))
	assert.Equal(t, []string{"A", "B", "D"}, exclude([]string{"A", "B", "D"}, "C"))
}
