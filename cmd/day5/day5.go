package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
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

func main() {
	input, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	reduced := reduce(input)

	fmt.Println(fmt.Sprintf("Part one: len(%v)", len(reduced)))
}
