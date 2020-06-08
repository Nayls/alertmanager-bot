package utils

import (
	"log"
	"net/url"
)

// GetParsedURL ...
func GetParsedURL(payload string) *url.URL {
	res, err := url.Parse(payload)
	if err != nil {
		log.Fatal("GetParseURL parse error: ", err)
		return nil
	}
	return res
}
