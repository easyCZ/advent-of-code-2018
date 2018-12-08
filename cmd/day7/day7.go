package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	stepRegex = regexp.MustCompile(`Step (?P<Step>\w) must be finished before step (?P<Dependency>\w) can begin.`)
)

func NewGraph(data map[string][]string) *Graph {
	return &Graph{
		data:      data,
		ancestors: reverseGraph(data),
	}
}

type Graph struct {
	data      map[string][]string
	ancestors map[string][]string
}

func (g *Graph) Add(step, dependsOn string) {
	deps, ok := g.data[step]
	if ok {
		deps = append(deps, dependsOn)
		sort.Strings(deps)
		g.data[step] = deps
	} else {
		g.data[step] = []string{dependsOn}
	}

	_, ok = g.data[dependsOn]
	if !ok {
		g.data[dependsOn] = []string{}
	}
}

func reverseGraph(g map[string][]string) map[string][]string {
	ancestor := make(map[string][]string)

	for k := range g {
		ancestor[k] = []string{}
	}

	for k, v := range g {
		for _, dep := range v {
			ancestors := append(ancestor[dep], k)
			ancestor[dep] = ancestors
		}
	}
	return ancestor
}

func (g *Graph) Root() []string {
	roots := make([]string, 0)
	for node, deps := range g.ancestors {
		if len(deps) == 0 {
			roots = append(roots, node)
		}
	}

	return roots
}

func (g *Graph) Order() []string {
	nextMove := func(visited map[string]bool, candidates []string) string {

		satisfiesDependencies := func(deps []string) bool {
			for _, dep := range deps {
				if !visited[dep] {
					return false
				}
			}
			return true
		}
		sort.Strings(candidates)
		for _, candidate := range candidates {
			// Does the candidate satisfy the dependencies
			if satisfiesDependencies(g.ancestors[candidate]) {
				return candidate
			}
		}

		panic("did not find a candidate")

		return ""
	}

	order := []string{}
	visited := make(map[string]bool)
	candidates := g.Root()

	for len(candidates) > 0 {
		next := nextMove(visited, candidates)
		visited[next] = true
		candidates = exclude(candidates, next)

		order = append(order, next)

		candidates = append(candidates, g.data[next]...)
	}

	return order
}

func ExecutePlan(g *Graph) {

}

func exclude(vals []string, val string) []string {
	out := make([]string, 0)
	for _, v := range vals {
		if v != val {
			out = append(out, v)
		}
	}
	return out
}

func (g *Graph) String() string {
	var b bytes.Buffer
	for k, v := range g.data {
		b.WriteString(fmt.Sprintf("%v: %v\n", k, v))
	}
	return b.String()
}

func parse(r io.Reader) *Graph {
	scanner := bufio.NewScanner(r)

	g := make(map[string][]string)

	add := func(step, dependsOn string) {
		deps, ok := g[step]
		if ok {
			deps = append(deps, dependsOn)
			sort.Strings(deps)
			g[step] = deps
		} else {
			g[step] = []string{dependsOn}
		}

		_, ok = g[dependsOn]
		if !ok {
			g[dependsOn] = []string{}
		}
	}

	for scanner.Scan() {
		t := scanner.Text()
		tokens := stepRegex.FindStringSubmatch(t)
		add(tokens[1], tokens[2])
	}

	return NewGraph(g)
}

func main() {
	graph := parse(os.Stdin)

	fmt.Println(fmt.Sprintf("Part one: %v", strings.Join(graph.Order(), "")))
}
