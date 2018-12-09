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
	stepRegex    = regexp.MustCompile(`Step (?P<Step>\w) must be finished before step (?P<Dependency>\w) can begin.`)
	baseStepCost = 60
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

func (g *Graph) Roots() []string {
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
		sort.Strings(candidates)
		for _, candidate := range candidates {
			// Does the candidate satisfy the dependencies
			if g.DependenciesSatisfied(candidate, visited) {
				return candidate
			}
		}

		panic("did not find a candidate")

		return ""
	}

	order := []string{}
	visited := make(map[string]bool)
	candidates := g.Roots()

	for len(candidates) > 0 {
		next := nextMove(visited, candidates)
		visited[next] = true
		candidates = exclude(candidates, next)

		order = append(order, next)

		candidates = append(candidates, g.data[next]...)
	}

	return order
}

func (g *Graph) DependenciesSatisfied(node string, visited map[string]bool) bool {
	for _, dep := range g.ancestors[node] {
		if !visited[dep] {
			return false
		}
	}
	return true
}

type Job struct {
	id        string
	completed int
	progress  int
}

func NewJob(id string) *Job {
	return &Job{
		id:        id,
		completed: stepCost(id),
		progress:  0,
	}
}

type Worker struct {
	job *Job
}

func (w *Worker) Tick() {
	if w.job != nil {
		w.job.progress += 1
	}
}

func (w *Worker) HasJob() bool {
	return w.job != nil
}

func (w *Worker) Done() bool {
	return w.job != nil && w.job.progress >= w.job.completed
}

func (w *Worker) Assign(j string) {
	w.job = NewJob(j)
}

func (w *Worker) Remaining() int {
	if w.HasJob() {
		return w.job.completed - w.job.progress
	}
	return 0
}

func stepCost(id string) int {
	return int(id[0]-"A"[0]) + baseStepCost + 1
}

func workersAllDone(workers []*Worker) bool {
	for _, w := range workers {
		if w.HasJob() {
			return false
		}
	}
	return true
}

func tick(workers []*Worker) {
	for _, w := range workers {
		w.Tick()
	}
}

func ExecutePlan(g *Graph, workerCount int) int {
	workers := make([]*Worker, 0)
	for i := 0; i < workerCount; i++ {
		workers = append(workers, &Worker{})
	}

	time := 0
	completed := make(map[string]bool)

	// determine initial work
	jobs := g.Roots()
	pendingJobs := make([]string, 0)

	// While there is work to be done
	for len(jobs) > 0 || len(pendingJobs) > 0 {
		// collect completed work
		for _, w := range workers {
			if w.HasJob() && w.Done() {
				completed[w.job.id] = true
				pendingJobs = exclude(pendingJobs, w.job.id)
				jobs = set(append(jobs, g.data[w.job.id]...))
				w.job = nil
			}
		}

		// attempt to assign work
		eligibleWorkers := make([]*Worker, 0)
		for _, w := range workers {
			if !w.HasJob() {
				eligibleWorkers = append(eligibleWorkers, w)
			}
		}

		eligibleJobs := make([]string, 0)
		for _, job := range jobs {
			if g.DependenciesSatisfied(job, completed) && !completed[job] {
				eligibleJobs = append(eligibleJobs, job)
			}
		}
		sort.Strings(eligibleJobs)

		allocs := 0
		for jobIndex, j := range eligibleJobs {
			if len(eligibleWorkers) > jobIndex {
				allocs += 1
				eligibleWorkers[jobIndex].Assign(j)
				pendingJobs = append(pendingJobs, j)
				jobs = exclude(jobs, j)
			}
		}

		// tick
		tick(workers)
		time += 1
	}

	return time - 1
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
	//	var b bytes.Buffer
	//	b.WriteString(`Step C must be finished before step A can begin.
	//Step C must be finished before step F can begin.
	//Step A must be finished before step B can begin.
	//Step A must be finished before step D can begin.
	//Step B must be finished before step E can begin.
	//Step D must be finished before step E can begin.
	//Step F must be finished before step E can begin.
	//`)
	graph := parse(os.Stdin)

	fmt.Println(fmt.Sprintf("Part one: %v", strings.Join(graph.Order(), "")))
	fmt.Println(fmt.Sprintf("Part two: %v", ExecutePlan(graph, 5)))
}

func set(vals []string) []string {
	index := make(map[string]bool)
	for _, val := range vals {
		index[val] = true
	}
	deduped := make([]string, 0)
	for k := range index {
		deduped = append(deduped, k)
	}
	return deduped
}
