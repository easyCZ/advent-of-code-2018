package main

import (
	"fmt"
	"strconv"
)

type Grid struct {
	vals [][]int
}

func hundredsDigit(num int) int {
	if num >= 100 {
		text := fmt.Sprintf("%d", num)
		n, err := strconv.ParseInt(string(text[len(text)-1-2]), 10, 32)
		if err != nil {
			return 0
		}
		return int(n)
	}
	return 0
}
func power(x, y, serial int) int {
	rackID := x + 10
	p := (rackID*y + serial) * rackID
	return hundredsDigit(p) - 5
}

func gridPower(grid [][]int) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			sum += grid[i][j]
		}
	}
	return sum
}

func newGrid(serial int) [][]int {
	grid := make([][]int, 300)
	for i := 0; i < 300; i++ {
		grid[i] = make([]int, 300)
	}

	for col := 0; col < 300; col++ {
		for row := 0; row < 300; row++ {
			grid[col][row] = power(col+1, row+1, serial)
		}
	}

	return grid
}

func maximiseSection(grid [][]int, size int) (x int, y int, maxPower int) {
	maxPower = 0
	maxX := 0
	maxY := 0

	for col := 0; col < len(grid)-size; col++ {
		for row := 0; row < len(grid[col])-size; row++ {
			section := make([][]int, size)
			for i := 0; i < size; i++ {
				section[i] = grid[col+i][row : row+size]
			}
			power := gridPower(section)
			if power > maxPower {
				maxX = col
				maxY = row
				maxPower = power
			}
		}
	}

	return maxX + 1, maxY + 1, maxPower
}

func maximiseSize(grid [][]int) (x, y, maxSize int) {
	maxSize = 0
	maxPower := 0
	maxX := 0
	maxY := 0
	for i := 1; i <= 300; i++ {
		x, y, power := maximiseSection(grid, i)
		if power > maxPower {
			maxSize = i
			maxPower = power
			maxX = x
			maxY = y
		}
	}

	return maxX + 1, maxY + 1, maxSize
}

func main() {
	grid := newGrid(2187)
	x, y, power := maximiseSection(grid, 3)
	fmt.Println(fmt.Sprintf("Part one: [%d, %d]  with power %d", x, y, power))

	x, y, size := maximiseSize(grid)
	fmt.Println(fmt.Sprintf("Part two: [%d, %d] with size %d", x, y, size))
}
