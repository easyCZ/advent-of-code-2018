package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

func parse(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	data := make([]string, 0)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if len(data) > 1 || len(data) == 0 {
		return "", errors.New("expected exactly one line in input")
	}

	return data[0], nil
}

func react(polymer string) string {
	for i := 0; i < len(polymer)-1; i++ {
		first := polymer[i]
		second := polymer[i+1]

		diff := int(first) - int(second)
		if diff == -32 || diff == 32 {
			return polymer[0:i] + polymer[i+2:]
		}
	}

	return polymer
}

func reduce(polymer string) string {
	reacted := react(polymer)
	if reacted == polymer || reacted == "" {
		return reacted
	}

	return reduce(reacted)
}

func elements(polymer string) []string {
	index := make(map[string]bool)

	for _, e := range polymer {
		index[strings.ToLower(string(e))] = true
	}

	elems := make([]string, 0)
	for elem := range index {
		elems = append(elems, elem)
	}

	return elems
}

func ShortestWithRemoval(polymer string) string {
	elems := elements(polymer)
	candidates := make([]string, 0)

	for _, elem := range elems {
		removed := removeCaseInsensitive(polymer, elem)
		candidates = append(candidates, reduce(removed))
	}

	shortest := candidates[0]
	for _, candidate := range candidates {
		if len(candidate) < len(shortest) {
			shortest = candidate
		}
	}

	return shortest
}

func removeCaseInsensitive(s string, val string) string {
	return strings.Replace(
		strings.Replace(s, strings.ToLower(val), "", -1),
		strings.ToUpper(val),
		"",
		-1,
	)
}

func main() {
	input, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	reduced := reduce(input)

	fmt.Println(fmt.Sprintf("Part one: len(%v)", len(reduced)))

	shortest := ShortestWithRemoval(input)
	fmt.Println(fmt.Sprintf("Part two: len(%v)", len(shortest)))
}
