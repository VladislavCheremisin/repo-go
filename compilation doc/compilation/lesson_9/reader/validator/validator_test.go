package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatorURL(t *testing.T) {
	input := "http://jaeger:16686"
	expected := "http://jaeger:16686"
	assert.Equal(t, input, expected, "they should be equal")
}

func TestNegativeValidatorURL(t *testing.T) {
	input := "http://jaeger:16686"
	expected := "xxxxxxxxxxxxx"
	assert.NotEqual(t, input, expected, "they should be equal")
}
