package main

import (
	"fmt"
	"os"

	"github.com/dishbreak/aoc2017/cmd/day15/part1"
	"github.com/dishbreak/aoc2017/cmd/day15/part2"
)

func main() {
	f, err := os.Open("inputs/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1.Solve(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2.Solve(f))
}
