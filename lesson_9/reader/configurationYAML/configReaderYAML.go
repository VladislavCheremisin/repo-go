package configurationReaderYAML

import (
	validationUrl "lesson_9/reader/urlValidator"
	"lesson_9/src/gopkg.in/yaml.v2"
	"os"
)

type ConfigurationYaml struct {
	Port        int    `yaml:"port"`
	DbURL       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppID   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

func ConfigReaderYAML() interface{} {
	data, err := os.ReadFile("configYAML.yaml")
	if err != nil {
		panic(err)
	}
	var parsed ConfigurationYaml
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		panic(err)
	}

	validationUrl.ValidatorUrl(parsed.DbURL)
	validationUrl.ValidatorUrl(parsed.JaegerURL)
	validationUrl.ValidatorUrl(parsed.SentryURL)

	return parsed
}
