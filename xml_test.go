package ydformatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-ydformatter"
)

func Test_FormatXML_Valid(t *testing.T) {

	source := "<root><data attrib=\"value\">data</data></root>"
	expected := `<root>
    <data attrib="value">data</data>
</root>`

	actual, err := ydformatter.XML(source)
	assert.NoError(t, err, "error")
	assert.Equal(t, expected, actual, "xml")

}

func Test_FormatXML_Invalid(t *testing.T) {

	source := "<root"

	actual, err := ydformatter.XML(source)
	assert.Error(t, err, "error")
	assert.Empty(t, actual, "xml")

}
