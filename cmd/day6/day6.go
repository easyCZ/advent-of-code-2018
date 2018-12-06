package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p *Point) Neighbours() []*Point {
	return []*Point{
		{x: p.x + 1, y: p.y},
		{x: p.x - 1, y: p.y},
		{x: p.x, y: p.y + 1},
		{x: p.x, y: p.y - 1},
		{x: p.x + 1, y: p.y + 1},
		{x: p.x - 1, y: p.y + 1},
		{x: p.x + 1, y: p.y - 1},
		{x: p.x + 1, y: p.y + 1},
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("[%d, %d]", p.x, p.y)
}

type Grid struct {
	points []*Point
}

func (g *Grid) Dimensions() (*Point, *Point) {
	var minX, minY, maxX, maxY int

	for _, p := range g.points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	return &Point{x: minX, y: minY}, &Point{x: maxX, y: maxY}
}

func (g *Grid) Width() int {
	topleft, botright := g.Dimensions()
	return botright.x - topleft.x
}

func (g *Grid) Height() int {
	topleft, botright := g.Dimensions()
	return botright.y - topleft.y
}

func (g *Grid) Plot() [][]*Point {
	width := g.Width()
	height := g.Height()

	plot := make([][]*Point, height)
	for i := 0; i < height; i++ {
		plot[i] = make([]*Point, width)
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			nearest, multiple := g.nearest(&Point{x: col, y: row})

			if multiple {
				plot[row][col] = nil
			} else {
				plot[row][col] = nearest
			}
		}
	}

	return plot
}

func (g *Grid) influence(p *Point, plot [][]*Point) int {
	width := g.Width() + 2
	height := g.Height() + 2
	topleft, _ := g.Dimensions()

	seen := make([][]bool, height)
	for i := 0; i < height; i++ {
		seen[i] = make([]bool, width)
	}

	neighbours := p.Neighbours()
	i := 0
	inf := 0
	for i < len(neighbours) {
		point := neighbours[i]
		i += 1

		x := point.x - topleft.x - 1
		y := point.y - topleft.y - 1
		if x < 0 || x > width-3 || y > height-3 || y < 0 {
			continue
		}

		if seen[y][x] {
			continue
		}

		plotted := plot[y][x]
		seen[y][x] = true

		if plotted == nil {
			continue
		}

		if plotted.x == p.x && plotted.y == p.y {
			inf += 1
			neighbours = append(neighbours, point.Neighbours()...)
		}
	}

	return inf
}

func (g *Grid) MaxInfluence() (*Point, int) {
	plot := g.Plot()
	// Remove points on the edges
	candidates := make(map[*Point]bool)
	for _, p := range g.points {
		candidates[p] = true
	}

	for i := 0; i < len(plot); i++ {
		delete(candidates, plot[i][0])
		delete(candidates, plot[i][len(plot[i])-1])
	}

	for i := 0; i < len(plot[0]); i++ {
		delete(candidates, plot[0][i])
		delete(candidates, plot[len(plot)-1][i])
	}

	var maxPoint *Point
	maxInfluence := math.MinInt32

	for p := range candidates {
		influence := g.influence(p, plot)
		if influence > maxInfluence {
			maxInfluence = influence
			maxPoint = p
		}
	}

	return maxPoint, maxInfluence
}

func (g *Grid) nearest(p *Point) (*Point, bool) {
	var nearest *Point
	distance := math.MaxInt32
	multiple := false

	for _, point := range g.points {
		d := ManhattanDistance(p, point)

		if d == distance {
			multiple = true
		}
		if d < distance {
			nearest = point
			distance = d
			multiple = false
		}
	}

	return nearest, multiple
}

func ManhattanDistance(a *Point, b *Point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func SumDistanceToAll(p *Point, points []*Point) int {
	s := 0
	for _, point := range points {
		s += ManhattanDistance(p, point)
	}

	return s
}

func (g *Grid) PointsWithinTotalDistance(limit int) []*Point {
	result := make([]*Point, 0)
	plot := g.Plot()

	for row := 0; row < len(plot); row++ {
		for col := 0; col < len(plot[0]); col++ {
			p := &Point{x: col, y: row}
			if SumDistanceToAll(p, g.points) < limit {
				result = append(result, p)
			}
		}
	}

	return result
}

func parse(r io.Reader) ([]*Point, error) {
	scanner := bufio.NewScanner(r)

	points := make([]*Point, 0)
	for scanner.Scan() {
		t := scanner.Text()
		tokens := strings.Split(t, ", ")
		if len(tokens) != 2 {
			return nil, errors.New(fmt.Sprintf("exactly 2 tokens expected, got %v", t))
		}

		x, err := strconv.ParseInt(tokens[0], 10, 32)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to parse x from %s", t))
		}
		y, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to parse x from %s", t))
		}

		points = append(points, &Point{
			x: int(x),
			y: int(y),
		})
	}

	return points, nil
}

func main() {
	points, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	grid := &Grid{points: points}
	maxPoint, maxInfluence := grid.MaxInfluence()
	fmt.Println(fmt.Sprintf("Part one: %v, influence: %v", maxPoint, maxInfluence))

	fmt.Println(fmt.Sprintf("Part two: %v", len(grid.PointsWithinTotalDistance(10000))))
}
