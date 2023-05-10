package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	vals := parse(input)
	fmt.Printf("Part 1: %d\n", part1(vals))
	fmt.Printf("Part 2: %d\n", part2(vals))
}

func parse(input []string) []int {
	parts := strings.Fields(input[0])

	result := make([]int, len(parts))
	for i, s := range parts {
		result[i], _ = strconv.Atoi(s)
	}

	return result
}

func part1(vals []int) int {
	m := make([]int, len(vals))
	copy(m, vals)

	seen := make(map[string]bool)

	steps := 0

	for ; ; steps++ {
		state := encode(m)

		if seen[state] {
			return steps
		}

		seen[state] = true

		maxVal, maxIdx := -1, -1

		for i, blk := range m {
			if blk > maxVal {
				maxVal = blk
				maxIdx = i
			}
		}

		m[maxIdx] = 0

		for i := 1; i <= maxVal; i++ {
			j := (maxIdx + i) % len(m)
			m[j]++
		}
	}
}

func part2(vals []int) int {
	m := make([]int, len(vals))
	copy(m, vals)

	seen := make(map[string]bool)

	steps := 0
	loopState := ""
	loopStart := -1

	for ; ; steps++ {
		state := encode(m)

		if seen[state] {
			if loopState == "" {
				loopState = state
				loopStart = steps
				seen = make(map[string]bool)
			} else {
				return steps - loopStart
			}
		}

		seen[state] = true

		maxVal, maxIdx := -1, -1

		for i, blk := range m {
			if blk > maxVal {
				maxVal = blk
				maxIdx = i
			}
		}

		m[maxIdx] = 0

		for i := 1; i <= maxVal; i++ {
			j := (maxIdx + i) % len(m)
			m[j]++
		}
	}
}

func encode(vals []int) string {
	s := make([]string, len(vals))
	for i, v := range vals {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}
