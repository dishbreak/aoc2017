package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaughtAt(t *testing.T) {
	type testCase struct {
		time, slot, depth int
		caught            bool
	}

	/*
		0: 3 // hit
		1: 2
		4: 4
		6: 4 // hit
	*/
	testCases := []testCase{
		{0, 0, 3, true},
		{0, 1, 2, false},
		{0, 4, 4, false},
		{0, 6, 4, true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.caught, caughtAt(tc.time, tc.slot, tc.depth))
		})
	}
}

func TestPart1(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`
	r := strings.NewReader(input)
	assert.Equal(t, 24, part1(r))
}

func TestPart2(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`
	r := strings.NewReader(input)
	assert.Equal(t, 10, part2(r))
}
