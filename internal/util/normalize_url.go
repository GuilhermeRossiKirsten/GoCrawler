package util

import (
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(parsedURL.String(), parsedURL.Scheme+"://"), nil
}