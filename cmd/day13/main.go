package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func caughtAt(time, slot, depth int) bool {
	return (time+slot)%((depth-1)*2) == 0
}

func part1(r io.Reader) int {
	acc := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		pts := strings.Split(l, ": ")
		slot, _ := strconv.Atoi(pts[0])
		depth, _ := strconv.Atoi(pts[1])

		if caughtAt(0, slot, depth) {
			acc += slot * depth
		}
	}

	return acc
}

func part2(r io.Reader) int {
	s := bufio.NewScanner(r)

	layers := make([]firewallLayer, 0)

	for s.Scan() {
		layers = append(layers, parseFirewallLayer(s.Text()))
	}

	checkDelay := func(delay int) bool {
		for _, layer := range layers {
			if layer.caughtAt(delay) {
				return false
			}
		}
		return true
	}

	for i := 0; true; i++ {
		if checkDelay(i) {
			return i
		}
	}

	return -1
}

type firewallLayer struct {
	slot, depth int
}

func (f firewallLayer) caughtAt(time int) bool {
	return (time+f.slot)%((f.depth-1)*2) == 0
}

func parseFirewallLayer(l string) firewallLayer {
	f := firewallLayer{}
	pts := strings.Split(l, ": ")
	f.slot, _ = strconv.Atoi(pts[0])
	f.depth, _ = strconv.Atoi(pts[1])

	return f
}
