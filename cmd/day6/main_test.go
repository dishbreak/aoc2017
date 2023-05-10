package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, part1([]int{0, 2, 7, 0}))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 4, part2([]int{0, 2, 7, 0}))
}
