// Package validator checks string for validity
//
// The ValidatorURL returns valid URL or err

package validator

import (
	"log"
	"net/url"
)

// ValidatorURL takes a string as input and checks if the url is valid
func ValidatorURL(f string) *url.URL {
	validURL, err := url.ParseRequestURI(f)
	if err != nil {
		log.Fatal(err)
	}
	return validURL
}
