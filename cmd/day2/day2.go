package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func read(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	input := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		input = append(input, t)
	}
	return input
}

func count(s string) map[int][]rune {
	counter := make(map[rune]int)
	for _, r := range s {
		val, ok := counter[r]
		if !ok {
			counter[r] = 1
		} else {
			counter[r] = val + 1
		}
	}

	reverse := make(map[int][]rune)
	for r, count := range counter {
		val, ok := reverse[count]
		if !ok {
			runes := make([]rune, 0)
			runes = append(runes, r)
			reverse[count] = runes
		} else {
			val = append(val, r)
			reverse[count] = val
		}
	}
	return reverse
}

func checksum(input []string) int {
	twos := 0
	threes := 0

	for _, s := range input {
		counts := count(s)
		_, ok := counts[2]
		if ok {
			twos += 1
		}
		_, ok = counts[3]
		if ok {
			threes += 1
		}

	}

	return twos * threes
}

func diff(a, b string) int {
	c := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			c += 1
		}
	}

	return c
}

func minDiff(in []string) (string, string) {
	var a string
	var b string
	min := math.MaxInt32

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			d := diff(in[i], in[j])
			if d < min {
				min = d
				a = in[i]
				b = in[j]
			}
		}
	}

	return a, b
}

func common(a, b string) string {
	var s string
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			s = fmt.Sprintf("%s%s", s, string(a[i]))
		}
	}
	return s
}

func main() {
	input := read(os.Stdin)
	fmt.Println("Part one: ", checksum(input))

	a, b := minDiff(input)
	fmt.Println("Part two: ", a, b)
	fmt.Println("Common  : ", fmt.Sprintf("%s", common(a, b)))
}
