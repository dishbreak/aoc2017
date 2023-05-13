package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day8.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type cpu struct {
	register map[string]int
	exprs    map[string]expr
	conds    map[string]cond
	highest  int
}

type cond func(string, int) bool
type expr func(string, int)

func (c *cpu) incr(reg string, arg int) {
	c.register[reg] += arg
}

func (c *cpu) decr(reg string, arg int) {
	c.register[reg] -= arg
}

func (c *cpu) gt(reg string, arg int) bool {
	return c.register[reg] > arg
}

func (c *cpu) gte(reg string, arg int) bool {
	return c.register[reg] >= arg
}

func (c *cpu) lt(reg string, arg int) bool {
	return c.register[reg] < arg
}

func (c *cpu) lte(reg string, arg int) bool {
	return c.register[reg] <= arg
}

func (c *cpu) eq(reg string, arg int) bool {
	return c.register[reg] == arg
}

func (c *cpu) neq(reg string, arg int) bool {
	return c.register[reg] != arg
}

func (c *cpu) evaluate(instr string) {
	pts := strings.Fields(instr)

	// c dec -10 if a >= 1
	arg, _ := strconv.Atoi(pts[6])
	if !c.conds[pts[5]](pts[4], arg) {
		return
	}

	arg, _ = strconv.Atoi(pts[2])
	c.exprs[pts[1]](pts[0], arg)

	if c.register[pts[0]] > c.highest {
		c.highest = c.register[pts[0]]
	}
}

func (c *cpu) largestVal() int {
	max := math.MinInt
	for _, v := range c.register {
		if v > max {
			max = v
		}
	}
	return max
}

func newCpu() *cpu {
	c := &cpu{
		register: make(map[string]int),
		highest:  math.MinInt,
	}

	c.conds = map[string]cond{
		">":  c.gt,
		">=": c.gte,
		"<":  c.lt,
		"<=": c.lte,
		"==": c.eq,
		"!=": c.neq,
	}

	c.exprs = map[string]expr{
		"inc": c.incr,
		"dec": c.decr,
	}

	return c
}

func part1(input []string) int {
	c := newCpu()
	for _, line := range input {
		if line == "" {
			continue
		}
		c.evaluate(line)
	}

	return c.largestVal()
}

func part2(input []string) int {
	c := newCpu()
	for _, line := range input {
		if line == "" {
			continue
		}
		c.evaluate(line)
	}

	return c.highest
}
