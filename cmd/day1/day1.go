package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
)


func read(r io.Reader) ([]int64,  error) {
	data := make([]int64, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()

		i, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return data, errors.Wrapf(err, "failed to parse %s", t)
		}
		data = append(data, i)
	}
	return data, nil
}

func sum(vals []int64) int64 {
	var s int64
	for _, v := range vals {
		s += v
	}
	return s
}

func firstRepeatedSum(vals []int64) int64 {
	seen := make(map[int64]bool)
	var currentSum int64
	index := 0
	for {
		curr := vals[index]
		currentSum += curr
		_, ok := seen[currentSum]
		if ok {
			return currentSum
		}
		seen[currentSum] = true
		index = (index + 1) % len(vals)
	}
}

func main() {
	data, err := read(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", sum(data))
	fmt.Println("Part 2:", firstRepeatedSum(data))
}
