package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []string{
		"5 1 9 5",
		"7 5 3",
		"2 4 6 8",
		"",
	}

	assert.Equal(t, 18, part1(input))
}

func TestPart2(t *testing.T) {
	input := []string{
		"5 9 2 8",
		"9 4 7 3",
		"3 8 6 5		",
		"",
	}

	assert.Equal(t, 9, part2(input))
}
