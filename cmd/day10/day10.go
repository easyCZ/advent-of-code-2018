package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	inputRegex = regexp.MustCompile(`position=<\s?(?P<X>-?\d+),\s*(?P<Y>-?\d+)> velocity=<\s*(?P<VelocityX>-?\d+),\s*(?P<VelocityY>-?\d+)>`)
)

type Vector struct {
	x int
	y int
}

type Point struct {
	position *Vector
	velocity *Vector
}

func (p *Point) String() string {
	return fmt.Sprintf(`[x: %v, y: %v] -> [x: %v, y: %v]`, p.position.x, p.position.y, p.velocity.x, p.velocity.y)
}

func (p *Point) Tick() {
	p.position = &Vector{
		x: p.position.x + p.velocity.x,
		y: p.position.y + p.velocity.y,
	}
}

type Grid struct {
	points []*Point
}

func (g *Grid) Tick() {
	for _, p := range g.points {
		p.Tick()
	}
}

func (g *Grid) Dimensions() (topleft *Vector, botright *Vector) {
	var minX, maxX int

	for _, p := range g.points {
		if p.position.x < minX {
			minX = p.position.x
		}
		if p.position.x > maxX {
			maxX = p.position.x
		}
	}

	var minY, maxY int
	for _, p := range g.points {
		if p.position.y < minY {
			minY = p.position.y
		}
		if p.position.y > maxY {
			maxY = p.position.y
		}
	}

	return &Vector{minX, minY}, &Vector{maxX, maxY}
}

func (g *Grid) Plot() [][]bool {
	topleft, botright := g.Dimensions()
	width := int(math.Abs(float64(botright.x - topleft.x)))
	height := int(math.Abs(float64(botright.y - topleft.y)))

	plot := make([][]bool, height)
	for i := 0; i < height; i++ {
		plot[i] = make([]bool, width)
	}

	for _, p := range g.points {
		col := int(math.Abs(float64(p.position.y + topleft.y)))
		row := int(math.Abs(float64(p.position.x + topleft.x)))
		if col < 0 || col > len(plot) {
			panic(fmt.Sprintf("col %v is out of bounds", col))
		}
		if row < 0 || row > len(plot[col]) {
			panic(fmt.Sprintf("row %v is out of bounds", row))
		}
		plot[col][row] = true
	}

	return plot
}

func (g *Grid) Draw() string {
	plot := g.Plot()

	var b bytes.Buffer
	for col := range plot {
		for row := range plot[col] {
			if plot[col][row] {
				b.WriteString("#")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

func parse(r io.Reader) ([]*Point, error) {
	//	position=<-3, 11> velocity=< 1, -2>
	scanner := bufio.NewScanner(r)

	points := make([]*Point, 0)
	for scanner.Scan() {
		t := scanner.Text()
		matches := inputRegex.FindStringSubmatch(t)
		x, err := strconv.ParseInt(matches[1], 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse %v", t)
		}

		y, err := strconv.ParseInt(matches[2], 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse %v", t)
		}

		velocityX, err := strconv.ParseInt(matches[3], 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse %v", t)
		}

		velocityY, err := strconv.ParseInt(matches[1], 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse %v", t)
		}

		points = append(points, &Point{
			position: &Vector{int(x), int(y)},
			velocity: &Vector{int(velocityX), int(velocityY)},
		})
	}

	return points, nil
}

func main() {
	points, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	grid := &Grid{points}

	for i := 0; i < 1; i++ {
		grid.Tick()
		fmt.Println()
		fmt.Println(grid.Draw())
	}
}
