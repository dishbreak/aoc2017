package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	acc := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		acc += pt1Checksum(line)
	}
	return acc
}

func part2(input []string) int {
	acc := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		acc += pt2Checksum(line)
	}

	return acc
}

func pt1Checksum(line string) int {
	min, max := 9999999, -1

	for _, v := range strings.Fields(line) {
		n, _ := strconv.Atoi(v)
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	return max - min
}

func pt2Checksum(line string) int {
	pts := strings.Fields(line)
	data := make([]int, len(pts))
	for i, val := range pts {
		data[i], _ = strconv.Atoi(val)
	}

	big, small := findModuloPair(data)
	return big / small
}

func findModuloPair(data []int) (int, int) {
	for i, a := range data {
		for j, b := range data {
			if i == j || a > b {
				continue
			}
			if b%a == 0 {
				return b, a
			}
		}
	}
	return -1, -1
}
