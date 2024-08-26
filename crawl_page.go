package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// Send an empty struct to the buffered channel
	// when a new goroutine starts
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		// Free up channel
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	// the current URL we're crawling
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing current URL")
		return
	}

	// Check if `rawCurrentUR` is on the same domain of `rawBaseURL`
	// If not, return the current `pages`
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalize current URL")
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrent)
	if !isFirst {
		return
	}

	fmt.Printf("crawling: %s\n", rawCurrentURL)

	// Get html body of the current URL
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error get HTML body:\n %v", err)
		return
	}

	// Get all the urls from the html body
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error get urls from html body:\n %v", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
