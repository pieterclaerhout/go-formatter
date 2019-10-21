package formatter_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func TestElapsed(t *testing.T) {

	type test struct {
		input    int64
		empty    string
		expected string
	}

	var tests = []test{
		{0, "-", "-"},
		{10, "-", "00:00:10"},
		{75, "-", "00:01:15"},
		{3675, "-", "01:01:15"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
			actual := formatter.Elapsed(tc.input, tc.empty)
			assert.Equal(t, actual, tc.expected)
		})
	}

}
