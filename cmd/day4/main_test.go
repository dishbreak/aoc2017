package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	valid bool
}

func TestPart1(t *testing.T) {

	testCases := []testCase{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd ee aa", false},
		{"aa bb cc dd ee aaa", true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.valid, ValidPassphrase(tc.input))
		})
	}
}

func TestPart2(t *testing.T) {
	/*
		abcde fghij is a valid passphrase.
		abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
		a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
		iiii oiii ooii oooi oooo is valid.
		oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
	*/
	testCases := []testCase{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"oiii ioii iioi iiio", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.valid, ValidPassphraseNoAnagrams(tc.input))
		})
	}
}
