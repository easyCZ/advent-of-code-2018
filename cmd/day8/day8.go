package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id       string
	children []*Node
	metadata []int
}

func (n *Node) SumMetadata() int {
	sum := 0
	for _, m := range n.metadata {
		sum += m
	}

	for _, c := range n.children {
		sum += c.SumMetadata()
	}

	return sum
}

func (n *Node) Value() int {
	if len(n.children) == 0 {
		return n.SumMetadata()
	}

	val := 0
	for _, md := range n.metadata {
		if md == 0 {
			continue
		}
		if len(n.children) > md-1 {
			val += n.children[md-1].Value()
		}
	}

	return val
}

func (n *Node) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(`
Node %v (
	metadata: %v,
	children: %s,
)`, n.id, n.metadata, n.children))
	return b.String()
}

func NewTree(input []int, id int) (*Node, []int) {
	node := &Node{
		id:       string(id + int("A"[0])),
		children: []*Node{},
		metadata: []int{},
	}
	childNodes := input[0]
	metadata := input[1]
	input = input[2:]

	for i := 0; i < childNodes; i++ {
		child, remainder := NewTree(input, id+1)
		node.children = append(node.children, child)
		input = remainder
	}

	node.metadata = input[:metadata]
	input = input[metadata:]

	return node, input
}

func parse(r io.Reader) (*Node, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	text := scanner.Text()
	tokens := strings.Split(text, " ")

	var input []int
	for _, t := range tokens {
		num, err := strconv.ParseInt(t, 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse %t", t)
		}
		input = append(input, int(num))
	}

	tree, _ := NewTree(input, 0)
	return tree, nil
}

func main() {
	root, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Part one: %v", root.SumMetadata()))
	fmt.Println(fmt.Sprintf("Part two: %v", root.Value()))
}
