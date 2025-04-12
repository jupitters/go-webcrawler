package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", errors.New("não foi possivel transformar em URL")
	}

	formatedURL := parsedURL.Host + strings.TrimRight(parsedURL.Path, "/")
	
	return formatedURL, nil
}