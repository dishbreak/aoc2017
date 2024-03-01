package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", score(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", countGarbage(f))
}

func score(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	return scoreRecursive(s, 0)
}

func scoreRecursive(s *bufio.Scanner, level int) int {
	acc := 0
	isGarbage := false
	canceled := false
	for s.Scan() {
		c := s.Text()
		if canceled {
			canceled = false
			continue
		}
		if c == "!" {
			canceled = true
			continue
		}
		if c == ">" {
			isGarbage = false
			continue
		}
		if isGarbage {
			continue
		}
		if c == "<" {
			isGarbage = true
		}
		switch c {
		case "{":
			acc += scoreRecursive(s, level+1)
		case "}":
			acc += level
			return acc
		}
	}
	return acc
}

func countGarbage(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)

	acc := 0
	garbage := false
	canceled := false
	for s.Scan() {
		c := s.Text()

		if canceled {
			canceled = false
			continue
		}

		if c == "!" {
			canceled = true
			continue
		}

		if c == ">" {
			garbage = false
			continue
		}

		if c == "<" && !garbage {
			garbage = true
			continue
		}
		if garbage {
			acc++
		}
	}
	return acc
}
