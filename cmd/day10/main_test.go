package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	input := []int{0, 1, 2, 3, 4}
	lengths := []int{3, 4, 1, 5}

	h := newKnotHasher(input)
	for _, length := range lengths {
		h.hash(length)
	}

	assert.Equal(t, []int{3, 4, 2, 1, 0}, h.data)
	assert.Equal(t, 4, h.pos)
	assert.Equal(t, 4, h.skip)
}
