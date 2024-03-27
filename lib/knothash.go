package lib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type KnotHasher struct {
	data []uint8
	pos  int
	skip int
}

func NewKnotHasher() *KnotHasher {
	d := make([]uint8, 256)
	for i := range d {
		d[i] = uint8(i)
	}

	return &KnotHasher{
		data: d,
		pos:  0,
		skip: 0,
	}
}

func (k *KnotHasher) hash(length uint8) {
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

func (k *KnotHasher) round(lengths []uint8) {
	for _, l := range lengths {
		k.hash(l)
	}
}

var suffix = []uint8{17, 31, 73, 47, 23}

const rounds = 64

func (k *KnotHasher) AddString(s string) {
	b := []byte(s)
	b = append(b, suffix...)

	for round := 0; round < rounds; round++ {
		k.round(b)
	}
}

func (k *KnotHasher) AddReader(r io.Reader) {
	b := make([]uint8, 0)

	s := bufio.NewReader(r)
	for bt, err := s.ReadByte(); err == nil; bt, err = s.ReadByte() {
		if bt == '\n' {
			continue
		}
		b = append(b, bt)
	}

	b = append(b, suffix...)
	for round := 0; round < rounds; round++ {
		k.round(b)
	}
}

func (k *KnotHasher) String() string {
	var sb strings.Builder

	for i := 0; i < 16; i++ {
		segment := k.data[i*16 : (i*16)+16]
		acc := segment[0]
		for j := 1; j < 16; j++ {
			acc ^= segment[j]
		}
		sb.WriteString(fmt.Sprintf("%02x", acc))
	}

	return sb.String()
}
