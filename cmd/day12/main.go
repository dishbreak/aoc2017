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
}

func part1(r io.Reader) int {
	g := parse(r)
	return g.trace(0)
}
