package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get CLI arguments
	args := os.Args[1:] // os.Args[0] is the program name, skip it
	baseURL := args[0]

	// Check the number of CLI arguments
	switch len(args) {
	case 0:
		fmt.Println("no website provided")
		return
	case 1:
		fmt.Printf("no maxConcurrency and maxPages provided")
	case 2:
		fmt.Printf("no maxPages provided")
	case 3:
		fmt.Printf("start crawling on: %v", baseURL)
	default:
		fmt.Println("too many arguments provided")
		return
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("failed to convert CLI arg to integer")
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("failed to convert CLI arg to integer")
	}

	cfg, err := createConfig(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("failed to create config struct:\n %v", err)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)

	fmt.Println("crawling ends")
}
