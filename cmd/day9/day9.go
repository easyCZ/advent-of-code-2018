package main

import (
	"container/list"
	"fmt"
)

type Ring struct {
	elems   *list.List
	current *list.Element
}

func (r *Ring) Place(marble int) {
	first := r.current.Next()
	if first == nil {
		first = r.elems.Front()
	}

	r.current = r.elems.InsertAfter(marble, first)
}

func (r *Ring) Remove() int {
	node := r.current
	for i := 0; i < 7; i++ {
		node = node.Prev()
		if node == nil {
			node = r.elems.Back()
		}
	}

	r.current = node.Next()
	if r.current == nil {
		r.current = r.elems.Front()
	}

	return r.elems.Remove(node).(int)
}

func (r *Ring) List() []int {
	node := r.elems.Front()
	items := make([]int, 0)
	for node != nil {
		items = append(items, node.Value.(int))
		node = node.Next()
	}

	return items
}

func NewRing() *Ring {
	elems := list.New()
	current := elems.PushFront(0)
	return &Ring{
		elems:   elems,
		current: current,
	}
}

type Game struct {
	ring       *Ring
	iterations int
	score      map[int]int
	player     int
	marble     int
}

func (g *Game) TakeTurn() {
	marble := g.marble + 1

	if marble%23 == 0 {
		score := marble + g.ring.Remove()
		g.score[g.player] = g.score[g.player] + score
	} else {
		g.ring.Place(marble)
	}

	g.marble = marble
	g.player = (g.player + 1) % len(g.score)
}

func (g *Game) Winner() (int, int) {
	max := -1
	maxId := -1
	for id, score := range g.score {
		if max < score {
			max = score
			maxId = id
		}
	}
	return maxId, max
}

func (g *Game) Play() int {
	for g.marble <= g.iterations {
		g.TakeTurn()
	}
	_, score := g.Winner()
	return score
}

func NewGame(players, max int) *Game {
	score := make(map[int]int)
	for i := 0; i < players; i++ {
		score[i] = 0
	}
	return &Game{
		ring:       NewRing(),
		iterations: max,
		player:     0,
		marble:     0,
		score:      score,
	}
}

func main() {
	game := NewGame(400, 71864)
	fmt.Println(fmt.Sprintf("Part one: %v", game.Play()))

	game2 := NewGame(400, 100 * 71864)
	fmt.Println(fmt.Sprintf("Part two: %v", game2.Play()))
}
