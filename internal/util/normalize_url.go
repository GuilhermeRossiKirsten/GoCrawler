package normalize_url

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(parsedURL.String(), parsedURL.Scheme+"://"), nil
}