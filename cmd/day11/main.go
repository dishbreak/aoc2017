package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// taken from https://www.redblobgames.com/grids/hexagons/
type point struct {
	q, r, s int
}

func pt(q, r, s int) point {
	return point{q, r, s}
}

func (p point) add(o point) (n point) {
	n.q = p.q + o.q
	n.r = p.r + o.r
	n.s = p.s + o.s
	return
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func dist(o, p point) (acc int) {
	acc += abs(o.q - p.q)
	acc += abs(o.r - p.r)
	acc += abs(o.s - p.s)
	acc /= 2
	return
}

var directions = map[string]point{
	"n":  pt(0, -1, 1),
	"nw": pt(-1, 0, 1),
	"sw": pt(-1, 1, 0),
	"s":  pt(0, 1, -1),
	"se": pt(1, 0, -1),
	"ne": pt(1, -1, 0),
}

type cursor struct {
	o point
}

func (c *cursor) move(s string) {
	dir, ok := directions[s]
	if !ok {
		panic(fmt.Errorf("unrecognized direction '%s'", s))
	}

	c.o = c.o.add(dir)
}

func main() {
	f, err := os.Open("inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rawData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if rawData[len(rawData)-1] == '\n' {
		rawData = rawData[:len(rawData)-1]
	}

	s := string(rawData)
	fmt.Printf("Part 1: %d\n", part1(s))
}

func part1(s string) int {
	steps := strings.Split(s, ",")

	c := &cursor{}

	for _, step := range steps {
		c.move(step)
	}

	return dist(c.o, point{})
}
