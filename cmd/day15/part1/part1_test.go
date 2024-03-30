package part1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `Generator A starts with 65
Generator B starts with 8921`
	r := strings.NewReader(input)

	assert.Equal(t, 588, Solve(r))
}
