package main

import (
	"fmt"

	"github.com/dishbreak/aoc2017/lib"
)

func main() {
	input := "ffayrhll"
	fmt.Printf("Part 1: %d\n", part1(input))
}

var parityBits = map[rune]int{
	'0': 0, // 0000
	'1': 1, // 0001
	'2': 1, // 0010
	'3': 2, // 0011
	'4': 1, // 0100
	'5': 2, // 0101
	'6': 2, // 0110
	'7': 3, // 0111
	'8': 1, // 1000
	'9': 2, // 1001
	'a': 2, // 1010
	'b': 3, // 1011
	'c': 2, // 1100
	'd': 3, // 1101
	'e': 3, // 1110
	'f': 4, // 1111
}

func part1(s string) int {
	acc := 0

	for i := 0; i < 128; i++ {
		k := lib.NewKnotHasher()
		key := fmt.Sprintf("%s-%d", s, i)
		k.AddString(key)
		for _, c := range k.String() {
			acc += parityBits[c]
		}
	}

	return acc
}
