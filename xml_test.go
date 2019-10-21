package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func TestFormatXMLValid(t *testing.T) {

	source := "<root><data attrib=\"value\">data</data></root>"
	expected := `<root>
    <data attrib="value">data</data>
</root>`

	actual, err := formatter.XML(source)
	assert.NoError(t, err, "error")
	assert.Equal(t, expected, actual, "xml")

}

func TestFormatXMLInvalid(t *testing.T) {

	source := "<root"

	actual, err := formatter.XML(source)
	assert.Error(t, err, "error")
	assert.Empty(t, actual, "xml")

}
