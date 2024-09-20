package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ABReverse(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(ABReverse(""), 0)
	ast.Equal(ABReverse("AAA"), 0)
	ast.Equal(ABReverse("BBB"), 0)
	ast.Equal(ABReverse("ABA"), 1)
	ast.Equal(ABReverse("ABAA"), 1)
	ast.Equal(ABReverse("ABAB"), 1)
	ast.Equal(ABReverse("ABABB"), 1)
	ast.Equal(ABReverse("ABABAA"), 2)
}
