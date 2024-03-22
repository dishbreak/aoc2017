package main

import (
	"fmt"
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

func TestXor(t *testing.T) {
	input := []uint8{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}
	acc := input[0]
	for j := 1; j < len(input); j++ {
		acc ^= input[j]
	}

	assert.Equal(t, uint8(64), acc)
	assert.Equal(t, "40", fmt.Sprintf("%02x", acc))
	assert.Equal(t, "05", fmt.Sprintf("%02x", uint8(5)))
}

func TestPart2(t *testing.T) {
	type testCase struct {
		input, result string
	}

	testCases := []testCase{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case", i), func(t *testing.T) {
			r := strings.NewReader(tc.input)
			assert.Equal(t, tc.result, part2(r))
		})
	}
}
