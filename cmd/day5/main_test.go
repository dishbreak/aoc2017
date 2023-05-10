package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, part1([]int{0, 3, 0, 1, -3}))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 10, part2([]int{0, 3, 0, 1, -3}))
}
