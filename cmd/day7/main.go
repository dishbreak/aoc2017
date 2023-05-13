package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day7.txt")
	if err != nil {
		panic(err)
	}

	input = input[:len(input)-1]

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) string {
	nodes := make(map[string]program)
	parent := make(map[string]string)

	for _, line := range input {
		p := ProgramFromLine(line)
		nodes[p.id] = p
		for _, dep := range p.deps {
			parent[dep] = p.id
		}
	}

	for n := range nodes {
		if _, ok := parent[n]; !ok {
			return n
		}
	}

	return "(not found)"
}

func part2(input []string) int {
	nodes := make(map[string]*program)
	parent := make(map[string]string)

	for _, line := range input {
		p := ProgramFromLine(line)
		nodes[p.id] = &p
		for _, dep := range p.deps {
			parent[dep] = p.id
		}
	}

	var start *program
	for _, n := range nodes {
		if pn, ok := parent[n.id]; !ok {
			start = nodes[n.id]
		} else {
			n.parent = nodes[pn]
		}
		for i, dep := range n.deps {
			n.next[i] = nodes[dep]
		}
	}

	sumUp(start)

	p := start
	for {
		var np *program
		for _, n := range p.next {
			if n.unbalanced {
				np = n
				break
			}
		}

		if np != nil {
			p = np
			continue
		}

		baseline := 0
		var trouble *program
		hits := make(map[int][]*program)

		for _, n := range p.next {
			hits[n.acc] = append(hits[n.acc], n)
		}

		for wt, matches := range hits {
			if len(matches) > 1 {
				baseline = wt
			} else {
				trouble = matches[0]
			}
		}

		return trouble.weight + baseline - trouble.acc
	}
}

func sumUp(p *program) int {
	p.acc = p.weight

	hits := make(map[int]int)
	for _, n := range p.next {
		w := sumUp(n)
		hits[w]++
		p.acc += w
	}

	if len(hits) != 1 {
		p.unbalanced = true
	}

	return p.acc
}
