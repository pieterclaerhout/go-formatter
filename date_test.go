package formatter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func TestUnixTimestamp(t *testing.T) {

	type test struct {
		input    int64
		layout   string
		expected string
	}

	var tests = []test{
		{1, "2006-01-02", "1970-01-01"},
		{1, "", ""},
		{-1000000, "2006-01-02", "1969-12-20"},
		{-1000000, "", ""},
	}

	for _, tc := range tests {
		name := fmt.Sprintf("%v", tc)
		t.Run(name, func(t *testing.T) {
			actual := formatter.UnixTimestamp(tc.input, tc.layout)
			assert.Equal(t, tc.expected, actual)
		})
	}

}

func TestDurationAsMilliseconds(t *testing.T) {

	type test struct {
		input    time.Duration
		expected string
	}

	var tests = []test{
		{1000000, "1ms"},
		{2000000, "2ms"},
		{2000005, "2ms"},
	}

	for _, tc := range tests {
		name := fmt.Sprintf("%v", tc)
		t.Run(name, func(t *testing.T) {
			actual := formatter.DurationInMilliseconds(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
