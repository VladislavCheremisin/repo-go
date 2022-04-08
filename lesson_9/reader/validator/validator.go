package validator

import (
	"log"
	"net/url"
)

func ValidatorURL(f string) *url.URL {
	validURL, err := url.ParseRequestURI(f)
	if err != nil {
		log.Fatal(err)
	}
	return validURL
}
