package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	acc := 0
	for _, pw := range input {
		if pw == "" {
			continue
		}
		if ValidPassphrase(pw) {
			acc++
		}
	}
	return acc
}

func part2(input []string) int {
	acc := 0
	for _, pw := range input {
		if pw == "" {
			continue
		}
		if ValidPassphraseNoAnagrams(pw) {
			acc++
		}
	}
	return acc
}

func ValidPassphrase(input string) bool {
	seen := make(map[string]bool)

	for _, part := range strings.Fields(input) {
		if seen[part] {
			return false
		}
		seen[part] = true
	}
	return true
}

func ValidPassphraseNoAnagrams(input string) bool {
	seen := make(map[string]bool)

	for _, part := range strings.Fields(input) {
		s := []rune(part)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		part = string(s)
		if seen[part] {
			return false
		}
		seen[part] = true
	}
	return true

}
