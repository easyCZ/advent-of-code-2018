package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tree = &Node{
		id:       "A",
		metadata: []int{1, 1, 2},
		children: []*Node{
			{id: "B", metadata: []int{10, 11, 12}, children: []*Node{}},
			{id: "C", metadata: []int{2}, children: []*Node{
				{id: "D", children: []*Node{}, metadata: []int{99}},
			}},
		},
	}
)

func TestNewTree(t *testing.T) {

	node, remainder := NewTree([]int{0, 1, 99, 2, 1, 1, 2}, 0)
	assert.Equal(t, remainder, []int{2, 1, 1, 2})
	assert.Len(t, node.children, 0)
	assert.Len(t, node.metadata, 1)

	node, remainder = NewTree([]int{1, 1, 0, 1, 99, 2, 1, 1, 2}, 0)
	assert.Equal(t, remainder, []int{1, 1, 2})
	assert.Len(t, node.children, 1)
	assert.Len(t, node.metadata, 1)
	assert.Equal(t, node.children[0].id, "B")

	node, remainder = NewTree([]int{0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}, 0)
	assert.Equal(t, remainder, []int{1, 1, 0, 1, 99, 2, 1, 1, 2})
	assert.Len(t, node.children, 0)
	assert.Len(t, node.metadata, 3)

	node, remainder = NewTree([]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}, 0)
	assert.Equal(t, remainder, []int{})
	assert.Len(t, node.children, 2)
	assert.Len(t, node.metadata, 3)

	assert.Len(t, node.children, 2)
	assert.Len(t, node.children[0].children, 0)
	assert.Len(t, node.children[1].children, 1)
	assert.Len(t, node.children[1].children[0].children, 0)
}

func TestNode_SumMetadata(t *testing.T) {
	assert.Equal(t, tree.SumMetadata(), 138)
}
