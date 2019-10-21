package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func TestColor(t *testing.T) {

	type test struct {
		input    string
		expected string
	}

	var tests = []test{
		{"aaaaaa", "AAAAAA"},
		{"bbbbbbb", "BBBBBB"},
		{"#cccccc", "CCCCCC"},
		{"#ddddddd", "DDDDDD"},
		{"eeeAAA", "EEEAAA"},
		{"aaaaCCC", "AAAACC"},
		{"#aaaBBB", "AAABBB"},
		{"#bbbAAAa", "BBBAAA"},
		{"#abc", "AABBCC"},
		{"#efe", "EEFFEE"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := formatter.Color(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
