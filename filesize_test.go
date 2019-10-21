package formatter_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func TestFileSize(t *testing.T) {

	type test struct {
		input    int64
		expected string
	}

	var tests = []test{
		{0, "0 bytes"},
		{100, "100 bytes"},
		{1000, "1.00 KB"},
		{1024, "1.02 KB"},
		{1100, "1.10 KB"},
		{2000, "2.00 KB"},
		{2048, "2.05 KB"},
		{1000000, "1.00 MB"},
		{1024000, "1.02 MB"},
		{1100000, "1.10 MB"},
		{2000000, "2.00 MB"},
		{2048000, "2.05 MB"},
		{1000000000, "1.00 GB"},
		{1024000000, "1.02 GB"},
		{1100000000, "1.10 GB"},
		{2000000000, "2.00 GB"},
		{2048000000, "2.05 GB"},
		{1000000000000, "1.00 TB"},
		{1024000000000, "1.02 TB"},
		{1100000000000, "1.10 TB"},
		{2000000000000, "2.00 TB"},
		{2048000000000, "2.05 TB"},
		{1000000000000000, "1000.00 TB"},
		{1024000000000000, "1024.00 TB"},
		{1100000000000000, "1100.00 TB"},
		{2000000000000000, "2000.00 TB"},
		{2048000000000000, "2048.00 TB"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
			actual := formatter.FileSize(tc.input)
			assert.Equal(t, actual, tc.expected)
		})
	}

}
