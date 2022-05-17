package reader

import (
	"os"
)

type Configuration struct {
	Port        int
	Db_url      string
	Jaeger_url  string
	Sentry_url  string
	KafkaBroker string
	SomeAppID   string
	SomeAppKey  string
}

func Reader(t string) []byte {

	if t == "yaml" {
		dataYaml, err := os.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		return dataYaml
	} else if t == "json" {
		dataJson, err := os.ReadFile("config.json")
		if err != nil {
			panic(err)
		}
		return dataJson
	}
	return nil
}
