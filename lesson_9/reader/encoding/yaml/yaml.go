package yaml

import (
	"lesson_9/reader/validator"
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

func ConfigReaderYaml() interface{} {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	var parsed ConfigurationYaml
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		panic(err)
	}

	validator.ValidatorURL(parsed.DbURL)
	validator.ValidatorURL(parsed.JaegerURL)
	validator.ValidatorURL(parsed.SentryURL)

	return parsed
}
