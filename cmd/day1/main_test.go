package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := map[string]int{
		"1122":     3,
		"1111":     4,
		"1234":     0,
		"91212129": 9,
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, expected, part1(input))
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]int{
		"1212":     6,
		"1221":     0,
		"123425":   4,
		"123123":   12,
		"12131415": 4,
	}

	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, expected, part2(input))
		})
	}
}
