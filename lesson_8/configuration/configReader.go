package configurationReader

import (
	"log"
	"net/url"
	"os"
)

type Configuration struct {
	Port         string
	Db_url       *url.URL
	Jaeger_url   *url.URL
	Sentry_url   *url.URL
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
}

func ConfigReader() interface{} {
	port := os.Getenv("PORT")
	db_url := os.Getenv("DB_URL")
	jaeger_url := os.Getenv("JAEGER_URL")
	sentry_url := os.Getenv("SENTRY_URL")
	kafka_broker := os.Getenv("KAFKA_BROKER")
	some_app_id := os.Getenv("SOME_APP_ID")
	some_app_key := os.Getenv("SOME_APP_KEY")

	jaeger_url_pars := validatorUrl(jaeger_url)
	sentry_url_pars := validatorUrl(sentry_url)
	db_url_pars := validatorUrl(db_url)

	validStruct := Configuration{
		Port:         port,
		Db_url:       db_url_pars,
		Jaeger_url:   jaeger_url_pars,
		Sentry_url:   sentry_url_pars,
		Kafka_broker: kafka_broker,
		Some_app_id:  some_app_id,
		Some_app_key: some_app_key,
	}
	return validStruct
}

func validatorUrl(f string) *url.URL {
	validURL, err := url.ParseRequestURI(f)
	if err != nil {
		log.Fatal(err)
	}
	return validURL
}
