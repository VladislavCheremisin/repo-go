package json

import (
	"encoding/json"
	"lesson_9/reader/validator"
	"os"
)

type ConfigurationJson struct {
	Port        int    `json:"port"`
	Db_url      string `json:"db_url"`
	Jaeger_url  string `json:"jaeger_url"`
	Sentry_url  string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppID   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func ConfigReaderJson() interface{} {
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var parsed ConfigurationJson
	if err := json.Unmarshal(data, &parsed); err != nil {
		panic(err)
	}

	validator.ValidatorURL(parsed.Db_url)
	validator.ValidatorURL(parsed.Jaeger_url)
	validator.ValidatorURL(parsed.Sentry_url)

	return parsed
}
