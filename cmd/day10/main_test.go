package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	input := []uint8{0, 1, 2, 3, 4}
	lengths := []uint8{3, 4, 1, 5}

	h := newKnotHasher(input)
	for _, length := range lengths {
		h.hash(length)
	}

	assert.Equal(t, []uint8{3, 4, 2, 1, 0}, h.data)
	assert.Equal(t, 4, h.pos)
	assert.Equal(t, 4, h.skip)
}

func TestProcessInput(t *testing.T) {
	input := "1,2,3\n"
	assert.Equal(t, []uint8{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}, processInput(strings.NewReader(input)))
}
