package validator

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"lesson_10/TestifyTests/reader/validator"
	"lesson_9/reader/reader"
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

func Example() {
	var parsed reader.Configuration
	config := reader.Reader("json")
	if err := json.Unmarshal(config, &parsed); err != nil {
		panic(err)
	}
	validator.ValidatorURL(parsed.Db_url)
}
