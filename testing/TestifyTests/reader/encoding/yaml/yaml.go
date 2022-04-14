package yaml

import (
	"lesson_9/reader/reader"
	"lesson_9/reader/validator"
	"lesson_9/src/gopkg.in/yaml.v2"
)

func ConfigReaderYaml() interface{} {

	var parsed reader.Configuration
	config := reader.Reader("yaml")
	if err := yaml.Unmarshal(config, &parsed); err != nil {
		panic(err)
	}

	validator.ValidatorURL(parsed.Db_url)
	validator.ValidatorURL(parsed.Jaeger_url)
	validator.ValidatorURL(parsed.Sentry_url)

	return parsed
}
