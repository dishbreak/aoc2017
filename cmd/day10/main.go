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
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
}

type knotHasher struct {
	data []int
	pos  int
	skip int
}

func newKnotHasher(d []int) *knotHasher {
	return &knotHasher{
		data: d,
		pos:  0,
		skip: 0,
	}
}

func (k *knotHasher) hash(length int) {
	var l []int
	wraparound := k.pos+length >= len(k.data)

	if !wraparound {
		l = k.data[k.pos : k.pos+length]
	} else {
		l = make([]int, length)
		for i := 0; i < length; i++ {
			l[i] = k.data[(k.pos+i)%len(k.data)]
		}
	}

	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}

	if wraparound {
		for i := 0; i < len(l); i++ {
			k.data[(k.pos+i)%len(k.data)] = l[i]
		}
	}

	k.pos = (k.pos + length + k.skip) % len(k.data)
	k.skip++
}

func extract(r io.Reader) []int {
	s := bufio.NewScanner(r)
	s.Scan()
	parts := strings.Split(s.Text(), ",")

	result := make([]int, len(parts))

	for i := range parts {
		result[i], _ = strconv.Atoi(parts[i])
	}

	return result
}

func part1(r io.Reader) int {
	lengths := extract(r)

	data := make([]int, 256)
	for i := range data {
		data[i] = i
	}

	h := newKnotHasher(data)

	for _, length := range lengths {
		h.hash(length)
	}

	return h.data[0] * h.data[1]
}

var suffix = []uint8{17, 31, 73, 47, 23}

func processInput(r io.Reader) []uint8 {
	result := make([]uint8, 0)

	s := bufio.NewReader(r)
	for b, err := s.ReadByte(); err == nil; b, err = s.ReadByte() {
		if b == '\n' {
			continue
		}
		result = append(result, b)
	}

	result = append(result, suffix...)
	return result
}
