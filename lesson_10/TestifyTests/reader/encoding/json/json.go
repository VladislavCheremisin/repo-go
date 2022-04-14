package json

import (
	"encoding/json"
	"lesson_9/reader/reader"
	"lesson_9/reader/validator"
)

func ConfigReaderJson() interface{} {

	var parsed reader.Configuration
	config := reader.Reader("json")
	if err := json.Unmarshal(config, &parsed); err != nil {
		panic(err)
	}

	validator.ValidatorURL(parsed.Db_url)
	validator.ValidatorURL(parsed.Jaeger_url)
	validator.ValidatorURL(parsed.Sentry_url)

	return parsed
}
