package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day15.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf("Part 1: %d\n", part1(f))
}

var factors = map[string]int{
	"A": 16807,
	"B": 48271,
}

const moduloValue = 2147483647

type generator struct {
	prev, factor int
}

var regexGeneratorLine = regexp.MustCompile(`^Generator (A|B) starts with (\d+)$`)

func parseLine(s string) *generator {
	matches := regexGeneratorLine.FindStringSubmatch(s)
	startVal, _ := strconv.Atoi(matches[2])
	factor := factors[matches[1]]

	return &generator{startVal, factor}
}

func (g *generator) generate() int {
	g.prev = (g.prev * g.factor) % moduloValue
	return g.prev
}

const mask = (1 << 16) - 1

func part1(r io.Reader) int {
	acc := 0

	s := bufio.NewScanner(r)
	s.Scan()
	genA := parseLine(s.Text())

	s.Scan()
	genB := parseLine(s.Text())

	for i := 0; i < 40000000; i++ {
		if a, b := genA.generate(), genB.generate(); (a & mask) == (b & mask) {
			acc++
		}
	}

	return acc
}
