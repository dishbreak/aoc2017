package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramFromString(t *testing.T) {
	type testCase struct {
		input  string
		result program
	}

	testCases := []testCase{
		{
			input: "pbga (66)",
			result: program{
				id:     "pbga",
				weight: 66,
			},
		},
		{
			input: "fwft (72) -> ktlj, cntj, xhth",
			result: program{
				id:     "fwft",
				weight: 72,
				deps: []string{
					"ktlj",
					"cntj",
					"xhth",
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, ProgramFromLine(tc.input))
		})
	}
}
