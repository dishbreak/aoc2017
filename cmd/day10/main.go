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
	data []uint8
	pos  int
	skip int
}

func newKnotHasher(d []uint8) *knotHasher {
	return &knotHasher{
		data: d,
		pos:  0,
		skip: 0,
	}
}

func (k *knotHasher) hash(length uint8) {
	var l []uint8
	wraparound := k.pos+int(length) >= len(k.data)

	if !wraparound {
		l = k.data[k.pos : k.pos+int(length)]
	} else {
		l = make([]uint8, length)
		for i := 0; i < int(length); i++ {
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

	k.pos = (k.pos + int(length) + k.skip) % len(k.data)
	k.skip++
}

func extract(r io.Reader) []uint8 {
	s := bufio.NewScanner(r)
	s.Scan()
	parts := strings.Split(s.Text(), ",")

	result := make([]uint8, len(parts))

	for i := range parts {
		tmp, _ := strconv.Atoi(parts[i])
		result[i] = uint8(tmp)
	}

	return result
}

func part1(r io.Reader) int {
	lengths := extract(r)

	data := make([]uint8, 256)
	for i := range data {
		data[i] = uint8(i)
	}

	h := newKnotHasher(data)

	for _, length := range lengths {
		h.hash(length)
	}

	return int(h.data[0]) * int(h.data[1])
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
