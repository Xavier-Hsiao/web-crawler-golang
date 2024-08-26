package main

import (
	"fmt"
	"os"
)

func main() {
	// Get CLI arguments
	args := os.Args[1:] // os.Args[0] is the program name, skip it

	// Check the number of CLI arguments
	switch len(args) {
	case 0:
		fmt.Println("no website provided")
		return
	case 1:
		fmt.Printf("start crawling of: %s\n", args[0])
	default:
		fmt.Println("too many arguments provided")
		return
	}

	const maxConcurrency = 2

	cfg, err := createConfig(args[0], 2)
	if err != nil {
		fmt.Printf("failed to create config struct:\n %v", err)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(args[0])
	cfg.wg.Wait()

	for url, count := range cfg.pages {
		fmt.Printf("%v - %v time(s)\n", url, count)
	}

	fmt.Println("crawling ends")
}
