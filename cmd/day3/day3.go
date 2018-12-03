package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	claims []*Claim
}

func (g *Grid) dimensions() (w int64, h int64) {
	maxX := int64(0)
	maxY := int64(0)

	for _, claim := range g.claims {
		if claim.left+claim.width > maxX {
			maxX = claim.left + claim.width
		}
		if claim.top+claim.height > maxY {
			maxY = claim.top + claim.height
		}
	}

	return maxX, maxY
}

func (g *Grid) plot() [][][]*Claim {
	width, height := g.dimensions()

	index := make([][][]*Claim, height)
	for i := range index {
		row := make([][]*Claim, width)
		for j := range row {
			row[j] = make([]*Claim, 0)
		}
		index[i] = row
	}

	for _, claim := range g.claims {
		for col := claim.top; col < claim.top+claim.height; col++ {
			for row := claim.left; row < claim.left+claim.width; row++ {
				index[col][row] = append(index[col][row], claim)
			}
		}
	}

	return index
}

func (g *Grid) countOverlap(minClaim int) int {
	c := 0
	plot := g.plot()
	for i := range plot {
		for j := range plot[i] {
			if len(plot[i][j]) >= minClaim {
				c += 1
			}
		}
	}

	return c
}

func (g *Grid) nonoverlapping() *Claim {
	plot := g.plot()
	for _, claim := range g.claims {
		eligible := true
		for col := claim.top; col < claim.top+claim.height && eligible; col++ {
			for row := claim.left; row < claim.left+claim.width && eligible; row++ {
				eligible = len(plot[col][row]) == 1 && eligible
			}
		}

		if eligible {
			return claim
		}
	}

	return nil
}

type Claim struct {
	id     string
	left   int64
	top    int64
	width  int64
	height int64
}

func tokenise(line string) (*Claim, error) {
	spaced := strings.Split(line, " ")

	id := strings.TrimLeft(spaced[0], "#")

	leftTop := strings.Split(strings.TrimRight(spaced[2], ":"), ",")
	left, err := strconv.ParseInt(leftTop[0], 10, 32)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse left from %s", line)
	}
	top, err := strconv.ParseInt(leftTop[1], 10, 32)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse top from %s", line)
	}

	size := strings.Split(spaced[3], "x")
	width, err := strconv.ParseInt(size[0], 10, 32)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse width from %s", line)
	}

	height, err := strconv.ParseInt(size[1], 10, 32)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse height from %s", line)
	}

	return &Claim{
		left:   left,
		height: height,
		top:    top,
		width:  width,
		id:     id,
	}, nil
}

func parse(r io.Reader) ([]*Claim, error) {
	scanner := bufio.NewScanner(r)

	claims := make([]*Claim, 0)
	for scanner.Scan() {
		claim, err := tokenise(scanner.Text())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to tokenise")
		}
		claims = append(claims, claim)
	}

	return claims, nil
}

func main() {
	claims, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	grid := &Grid{claims}
	fmt.Println(fmt.Sprintf("Part one: %d", grid.countOverlap(2)))
	fmt.Println(fmt.Sprintf("Part two: %s", grid.nonoverlapping().id))
}
