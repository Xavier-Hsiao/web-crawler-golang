package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("network error: %v", err)
	}
	defer res.Body.Close()

	// return 400+ status code
	if res.StatusCode > 399 {
		return "", fmt.Errorf("bad status code: %v: ", res.StatusCode)
	}

	// handle wrong res header content type
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("content type is not text/html")
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
