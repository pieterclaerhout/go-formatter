package ydformatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-ydformatter"
)

func Test_JSON_Valid(t *testing.T) {
	actual, err := ydformatter.JSONString("{\"key\": 1}")
	assert.NotEmpty(t, actual)
	assert.NoError(t, err)
}

func Test_JSON_Invalid(t *testing.T) {
	expected := "{\"key\": 1"
	actual, err := ydformatter.JSONString(expected)
	assert.Empty(t, actual)
	assert.Error(t, err)
}
