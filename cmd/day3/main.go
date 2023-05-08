package main

import (
	"fmt"
	"image"
)

const input = 325489

func main() {
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

var (
	north     = image.Pt(0, -1)
	south     = image.Pt(0, 1)
	east      = image.Pt(1, 0)
	west      = image.Pt(-1, 0)
	northeast = north.Add(east)
	northwest = north.Add(west)
	southeast = south.Add(east)
	southwest = south.Add(west)
)

var neighborhood = []image.Point{
	north, south, east, west, northeast, northwest, southeast, southwest,
}

func part1(input int) int {
	ring := 0
	maxVal := 0
	var pt image.Point
	for {
		b := 2*ring + 1
		maxVal = b * b

		if maxVal == input {
			return b * 2
		}

		if maxVal > input {
			pt = image.Pt(ring, ring)
			break
		}

		ring++
	}

	sideLen := ring * 2
	corners := make([]int, 3)
	for i := 0; i < 3; i++ {
		corners[i] = maxVal - sideLen*(i+1)
	}

	dirs := []image.Point{north, east, south}
	dir := west

	n := maxVal

	for {
		if n == input {
			return dist(pt)
		}

		if n == corners[0] {
			corners = corners[1:]
			dir = dirs[0]
			dirs = dirs[1:]
		}
		n--
		pt = pt.Add(dir)
	}

}

func dist(p image.Point) int {
	return abs(p.X) + abs(p.Y)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func part2(input int) int {
	m := make(map[image.Point]int)

	m[image.Pt(0, 0)] = 1
	ring := 1
	pt := image.Pt(1, 0)

	for {
		corners := []image.Point{
			image.Pt(ring, -1*ring),
			image.Pt(-1*ring, -1*ring),
			image.Pt(-1*ring, ring),
		}
		dirs := []image.Point{
			west,
			south,
			east,
		}

		dir := north

		for {
			acc := 0
			for _, n := range neighborhood {
				acc += m[pt.Add(n)]
			}
			if acc > input {
				return acc
			}
			m[pt] = acc
			if len(corners) > 0 && pt.Eq(corners[0]) {
				corners = corners[1:]
				dir = dirs[0]
				dirs = dirs[1:]
			}
			if pt.Eq(image.Pt(ring, ring)) {
				pt = pt.Add(east)
				ring++
				break
			}
			pt = pt.Add(dir)
		}
	}
}
