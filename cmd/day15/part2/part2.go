package part2

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

const moduloValue = 2147483647

type generator struct {
	prev, factor, multipleCheck int
}

var factors = map[string]int{
	"A": 16807,
	"B": 48271,
}

var multipleCheck = map[string]int{
	"A": 4,
	"B": 8,
}

var regexGeneratorLine = regexp.MustCompile(`^Generator (A|B) starts with (\d+)$`)

func parseLine(s string) *generator {
	matches := regexGeneratorLine.FindStringSubmatch(s)
	startVal, _ := strconv.Atoi(matches[2])
	factor := factors[matches[1]]
	mc := multipleCheck[matches[1]]

	return &generator{startVal, factor, mc}
}

func (g *generator) generate() int {
	g.prev = (g.prev * g.factor) % moduloValue

	for g.prev%g.multipleCheck != 0 {
		g.prev = (g.prev * g.factor) % moduloValue
	}

	return g.prev
}

const mask = (1 << 16) - 1

func Solve(r io.Reader) int {
	acc := 0

	s := bufio.NewScanner(r)
	s.Scan()
	genA := parseLine(s.Text())

	s.Scan()
	genB := parseLine(s.Text())

	for i := 0; i < 5000000; i++ {
		if a, b := genA.generate(), genB.generate(); (a & mask) == (b & mask) {
			acc++
		}
	}

	return acc
}
