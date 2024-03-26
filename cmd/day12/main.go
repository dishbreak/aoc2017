package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	g := parse(r)
	return g.trace(0)
}

func part2(r io.Reader) int {
	g := parse(r)
	acc := 0

	for id := range g {
		if g.trace(id) != -1 {
			acc++
		}
	}

	return acc
}
