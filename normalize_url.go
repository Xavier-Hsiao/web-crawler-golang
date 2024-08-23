package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(originalURL string) (string, error) {
	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse url")
	}

	fullPath := parsedURL.Hostname() + parsedURL.Path
	fullPath = strings.ToLower(fullPath)

	// Take off the trailing slash
	normalized := strings.TrimSuffix(fullPath, "/")
	return normalized, nil
}
