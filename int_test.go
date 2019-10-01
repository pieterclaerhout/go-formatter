package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func Test_IntWithSeparators(t *testing.T) {

	type test struct {
		name   string
		input  int64
		output string
	}

	var tests = []test{
		{"0", 0, "0"},
		{"1", 1, "1"},
		{"10", 10, "10"},
		{"100", 100, "100"},
		{"1000", 1000, "1.000"},
		{"10000", 10000, "10.000"},
		{"100000", 100000, "100.000"},
		{"1000000", 1000000, "1.000.000"},
		{"10000000", 10000000, "10.000.000"},
		{"100000000", 100000000, "100.000.000"},
		{"-0", -0, "0"},
		{"-1", -1, "-1"},
		{"-10", -10, "-10"},
		{"-100", -100, "-100"},
		{"-1000", -1000, "-1.000"},
		{"-10000", -10000, "-10.000"},
		{"-100000", -100000, "-100.000"},
		{"-1000000", -1000000, "-1.000.000"},
		{"-10000000", -10000000, "-10.000.000"},
		{"-100000000", -100000000, "-100.000.000"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := formatter.IntWithSeparators(tc.input)
			assert.Equal(t, tc.output, actual, tc.name)
		})
	}

}
