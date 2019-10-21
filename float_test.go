package formatter_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func TestFloat(t *testing.T) {

	type test struct {
		input     float64
		precision int64
		expected  string
	}

	var tests = []test{
		{0, 2, "0.00"},
		{0.1, 2, "0.10"},
		{0.11, 2, "0.11"},
		{0.111, 2, "0.11"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%f", tc.input), func(t *testing.T) {
			actual := formatter.FloatWithPrecision(tc.input, tc.precision)
			assert.Equal(t, actual, tc.expected)
		})
	}

}
