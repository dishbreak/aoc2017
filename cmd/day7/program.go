package main

import (
	"regexp"
	"strconv"
	"strings"
)

type program struct {
	id         string
	weight     int
	deps       []string
	acc        int
	next       []*program
	unbalanced bool
}

var prefixExp = regexp.MustCompile(`(\w+) \((\d+)\)`)

func ProgramFromLine(input string) program {
	parts := strings.Split(input, " -> ")
	p := program{}
	matches := prefixExp.FindStringSubmatch(parts[0])
	p.id = matches[1]
	p.weight, _ = strconv.Atoi(matches[2])

	if len(parts) > 1 {
		p.deps = strings.Split(parts[1], ", ")
	}

	p.next = make([]*program, len(p.deps))

	return p
}
