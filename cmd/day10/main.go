package main

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
