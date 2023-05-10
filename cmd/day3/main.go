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

/*
	you can consider the spiral as a series of rings.

	3333333
	3222223
	3211123
	3210123
	3211123
	3222223
	3333333

	the lower right-hand corner always has the highest address for the ring, like so

	3333333
	3222223
	3211123
	321*123
	3211*23
	32222*3
	333333*

	Additionally, the corner value is always the square of the nth odd number, where n is the ring number.
	1^2 = 1
	3^2 = 9
	5^2 = 25
	7^2 = 49
*/

func part1(input int) int {
	ring := 0
	maxVal := 0
	var pt image.Point

	// use a greedy algorithm to find the ring that the input falls on.
	// this will help us determine the starting point for our traversal.
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

	// the ring number tells us how long a side of the ring is.
	sideLen := ring * 2

	// compute the addresses for the other 3 corners of the ring
	// we will change traversal direction when we hit a corner value.
	corners := make([]int, 3)
	for i := 0; i < 3; i++ {
		corners[i] = maxVal - sideLen*(i+1)
	}

	// set the direction after we hit the corner.
	dirs := []image.Point{north, east, south}
	dir := west

	n := maxVal

	// we'll start clockwise from the maximum value
	for {
		// if we've found the target value, return the manhattan distance of the point.
		if n == input {
			return dist(pt)
		}

		// if we hit the next corner...
		if n == corners[0] {
			// pop the corner off the queue
			corners = corners[1:]
			// and change our direction.
			dir = dirs[0]
			dirs = dirs[1:]
		}
		// decrement our address by one
		n--
		// move one step in the given direction.
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
	/*
		we're going to start in ring 1, at the lowest address.
		We'll then traverse each ring, collecting neighbor sums as we go.
	*/
	m := make(map[image.Point]int)

	m[image.Pt(0, 0)] = 1
	ring := 1
	pt := image.Pt(1, 0)

	// iterate through progressive rings.
	for {
		// for each ring, identify the 3 turns that we'll take.
		corners := []image.Point{
			image.Pt(ring, -1*ring),
			image.Pt(-1*ring, -1*ring),
			image.Pt(-1*ring, ring),
		}
		// define the bearing for each corner.
		dirs := []image.Point{
			west,
			south,
			east,
		}

		// we'll start by heading north.
		dir := north

		// traverse each address in the ring.
		for {
			// set up an accumulator
			acc := 0

			// we're taking advantage of the zero value i.e. non-present keys will return 0
			for _, n := range neighborhood {
				acc += m[pt.Add(n)]
			}

			// if we've hit our target, return
			if acc > input {
				return acc
			}

			// otherwise, save the value to the map
			m[pt] = acc

			// check if we've hit a corner and turn if needed.
			if len(corners) > 0 && pt.Eq(corners[0]) {
				corners = corners[1:]
				dir = dirs[0]
				dirs = dirs[1:]
			}

			// if we've hit the last point in the ring, move the traversal into the next ring and stop iterating.
			if pt.Eq(image.Pt(ring, ring)) {
				pt = pt.Add(east)
				ring++
				break
			}

			// progress to the next point.
			pt = pt.Add(dir)
		}
	}
}
