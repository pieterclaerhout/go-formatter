package ydformatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-ydformatter"
)

func Test_Color(t *testing.T) {

	type test struct {
		input    string
		expected string
	}

	var tests = []test{
		{"aaaaaa", "AAAAAA"},
		{"aaaaaaa", "AAAAAA"},
		{"#aaaaaa", "AAAAAA"},
		{"#aaaaaaa", "AAAAAA"},
		{"aaaAAA", "AAAAAA"},
		{"aaaaAAA", "AAAAAA"},
		{"#aaaAAA", "AAAAAA"},
		{"#aaaAAAa", "AAAAAA"},
		{"#abc", "AABBCC"},
		{"#efe", "EEFFEE"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := ydformatter.Color(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
