package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	wg                 *sync.WaitGroup
	concurrencyControl chan struct{} // Buffered channel, avoiding spawn too many go routines
}

func createConfig(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL")
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
	}, nil
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	// Check if the current URL has been visited
	// If so, increment 1
	if count, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL] = count + 1
		return false
	}

	// If not, create a new entry
	cfg.pages[normalizedURL] = 1
	return true
}
