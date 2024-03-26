package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type node struct {
	id        int
	grouped   bool
	neighbors map[int]*node
}

type graph map[int]*node

func newNode(id int) *node {
	return &node{
		id:        id,
		neighbors: make(map[int]*node),
	}
}

func (g graph) link(one, other int) {
	oneNode, ok := g[one]
	if !ok {
		oneNode = newNode(one)
		g[one] = oneNode
	}

	otherNode, ok := g[other]
	if !ok {
		otherNode = newNode(other)
		g[other] = otherNode
	}

	oneNode.neighbors[other] = otherNode
	otherNode.neighbors[one] = oneNode
}

func parse(r io.Reader) graph {
	result := graph(make(map[int]*node))

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		pts := strings.Split(line, " <-> ")

		one, _ := strconv.Atoi(pts[0])

		others := strings.Split(pts[1], ", ")
		for _, other := range others {
			id, _ := strconv.Atoi(other)
			result.link(one, id)
		}
	}

	return result
}

func (g graph) trace(startingId int) int {
	visited := make(map[int]bool)
	start := g[startingId]
	if start.grouped {
		return -1
	}

	q := []*node{start}

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		if visited[n.id] {
			continue
		}

		visited[n.id] = true
		for _, p := range n.neighbors {
			q = append(q, p)
		}
	}

	for id, val := range visited {
		g[id].grouped = val
	}

	return len(visited)
}
