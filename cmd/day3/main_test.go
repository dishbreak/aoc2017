package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input, result int
}

func TestPart1(t *testing.T) {

	testCases := []testCase{
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, part1(tc.input))
		})
	}
}

func TestPart2(t *testing.T) {
	/*
		147  142  133  122   59
		304    5    4    2   57
		330   10    1    1   54
		351   11   23   25   26
		362  747  806--->   ...
	*/

	testCases := []testCase{
		{24, 25},
		{305, 330},
		{500, 747},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, part2(tc.input))
		})
	}
}
