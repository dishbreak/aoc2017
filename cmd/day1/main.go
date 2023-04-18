package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
	fmt.Printf("Part 2: %d\n", part2(input[0]))
}

func part1(input string) (checksum int) {
	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			checksum += int(input[i] - '0')
		}
	}

	if input[0] == input[len(input)-1] {
		checksum += int(input[0] - '0')
	}

	return
}

func part2(input string) (checksum int) {
	shift := len(input) / 2
	for i := range input {
		n := (i + shift) % len(input)
		if input[i] == input[n] {
			checksum += int(input[i] - '0')
		}
	}
	return
}
