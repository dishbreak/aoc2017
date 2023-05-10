package main

import (
	"fmt"
	"strconv"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	vals := parse(input)
	fmt.Printf("Part 1: %d\n", part1(vals))
	vals = parse(input)
	fmt.Printf("Part 2: %d\n", part2(vals))
}

func part1(vals []int) int {

	steps := 0
	for idx := 0; idx >= 0 && idx < len(vals); steps++ {
		offset := vals[idx]
		vals[idx]++
		idx += offset
	}

	return steps
}

func part2(vals []int) int {

	steps := 0
	for idx := 0; idx >= 0 && idx < len(vals); steps++ {
		offset := vals[idx]
		if offset >= 3 {
			vals[idx]--
		} else {
			vals[idx]++
		}
		idx += offset
	}

	return steps
}

func parse(input []string) []int {
	input = input[:len(input)-1]
	result := make([]int, len(input))

	for i, s := range input {
		result[i], _ = strconv.Atoi(s)
	}

	return result
}
