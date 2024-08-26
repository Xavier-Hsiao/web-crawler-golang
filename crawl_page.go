package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// base URL is the root URL we're crawling
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL")
		return
	}

	// the current URL we're crawling
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing current URL")
		return
	}

	// Check if `rawCurrentUR` is on the same domain of `rawBaseURL`
	// If not, return the current `pages`
	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalize current URL")
		return
	}

	// Check if the current URL has been visited
	// If so, increment 1
	if count, exists := pages[normalizedCurrent]; exists {
		pages[normalizedCurrent] = count + 1
		return
	}

	// If not, create a new entry
	pages[normalizedCurrent] = 1

	fmt.Printf("crawling: %s\n", rawCurrentURL)

	// Get html body of the current URL
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error get HTML body:\n %v", err)
		return
	}

	// Get all the urls from the html body
	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error get urls from html body:\n %v", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
