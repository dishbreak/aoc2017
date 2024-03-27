package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 8108, part1("flqrgnkx"))
}

func TestParityBits(t *testing.T) {
	input := "a0c2017"
	result := 9 // 10100000110000100000000101110000

	actual := 0

	for _, c := range input {
		actual += parityBits[c]
	}

	assert.Equal(t, result, actual)
}
