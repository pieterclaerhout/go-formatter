package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func TestJSONValid(t *testing.T) {
	actual, err := formatter.JSONString("{\"key\": 1}")
	assert.NotEmpty(t, actual)
	assert.NoError(t, err)
}

func TestJSONInvalid(t *testing.T) {
	expected := "{\"key\": 1"
	actual, err := formatter.JSONString(expected)
	assert.Empty(t, actual)
	assert.Error(t, err)
}
