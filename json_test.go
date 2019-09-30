package ydformatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-ydformatter"
)

func Test_JSON_Valid(t *testing.T) {
	actual := ydformatter.JSONString("{\"key\": 1}")
	assert.NotEmpty(t, actual)
}

func Test_JSON_Invalid(t *testing.T) {
	expected := "{\"key\": 1"
	actual := ydformatter.JSONString(expected)
	assert.Equal(t, expected, actual)
}
