package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day16.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %s\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %s\n", part2(f))
}

var (
	reSpin     = regexp.MustCompile(`s(\d+)`)
	reExchange = regexp.MustCompile(`x(\d+)/(\d+)`)
	reSwap     = regexp.MustCompile(`p(\w)/(\w)`)
)

func Spin(p []byte, i int) (b []byte) {
	i = i % len(p)
	b = append(b, p[len(p)-i:]...)
	b = append(b, p[0:len(p)-i]...)
	return b
}

func Exchange(p []byte, i, j int) (b []byte) {
	b = make([]byte, len(p))
	copy(b, p)
	b[i], b[j] = b[j], b[i]
	return
}

func Swap(p []byte, one, other byte) []byte {
	i, j := bytes.IndexByte(p, one), bytes.IndexByte(p, other)
	return Exchange(p, i, j)
}

func Parse(p []byte, inst string) []byte {

	if m := reSpin.FindStringSubmatch(inst); m != nil {
		l, _ := strconv.Atoi(m[1])
		return Spin(p, l)
	}

	if m := reExchange.FindStringSubmatch(inst); m != nil {
		i, _ := strconv.Atoi(m[1])
		j, _ := strconv.Atoi(m[2])
		return Exchange(p, i, j)
	}

	if m := reSwap.FindStringSubmatch(inst); m != nil {
		a := byte(m[1][0])
		b := byte(m[2][0])
		return Swap(p, a, b)
	}

	panic(fmt.Errorf("unrecognized instruction %s", inst))
}

func part1(r io.Reader) string {
	p := []byte("abcdefghijklmnop")

	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	insts := strings.Split(string(data), ",")
	for _, inst := range insts {
		p = Parse(p, inst)
	}

	return string(p)
}

const start = "abcdefghijklmnop"

func part2(r io.Reader) string {
	p := []byte(start)

	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	insts := strings.Split(string(data), ",")

	cycles := 0
	for i := 0; i < 1000000000; i++ {

		for _, inst := range insts {
			p = Parse(p, inst)
		}

		if string(p) == start {
			cycles = i + 1
			break
		}
	}

	for i := 0; i < (1000000000 % cycles); i++ {
		for _, inst := range insts {
			p = Parse(p, inst)
		}
	}

	return string(p)
}
