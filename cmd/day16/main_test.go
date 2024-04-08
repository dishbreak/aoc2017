package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpin(t *testing.T) {
	type testCase struct {
		input  string
		amount int
		result string
	}

	testCases := map[string]testCase{
		"base path": {
			"abcde", 1, "eabcd",
		},
		"wraparound with modulo": {
			"abcde", 6, "eabcd",
		},
		"spinning by length has no effect": {
			"abcde", 5, "abcde",
		},
		"spinning by zero has no effect": {
			"abcde", 0, "abcde",
		},
		"spinning by 3": {
			"abcde", 3, "cdeab",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			p := []byte(tc.input)
			p = Spin(p, tc.amount)
			assert.Equal(t, tc.result, string(p))
		})
	}
}

func TestParse(t *testing.T) {
	type testCase struct {
		input, inst, result string
	}

	testCases := map[string]testCase{
		"s1, a spin of size 1: eabcd.":                 {"abcde", "s1", "eabcd"},
		"x3/4, swapping the last two programs: eabdc.": {"eabcd", "x3/4", "eabdc"},
		"pe/b, swapping programs e and b: baedc.":      {"eabdc", "pe/b", "baedc"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			p := []byte(tc.input)
			p = Parse(p, tc.inst)
			assert.Equal(t, tc.result, string(p))
		})
	}
}
