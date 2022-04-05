package validationUrl

import (
	"log"
	"net/url"
)

func ValidatorUrl(f string) *url.URL {
	validURL, err := url.ParseRequestURI(f)
	if err != nil {
		log.Fatal(err)
	}
	return validURL
}
