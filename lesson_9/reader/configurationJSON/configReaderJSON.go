package configurationReaderJSON

import (
	"encoding/json"
	validationUrl "lesson_9/reader/urlValidator"
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

func ConfigReaderJSON() interface{} {
	data, err := os.ReadFile("configJSON.json")
	if err != nil {
		panic(err)
	}
	var parsed ConfigurationJson
	if err := json.Unmarshal(data, &parsed); err != nil {
		panic(err)
	}

	validationUrl.ValidatorUrl(parsed.Db_url)
	validationUrl.ValidatorUrl(parsed.Jaeger_url)
	validationUrl.ValidatorUrl(parsed.Sentry_url)

	return parsed
}
