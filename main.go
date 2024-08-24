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
		os.Exit(1)
	case 1:
		fmt.Printf("starting crawl of: %s", args[0])
	default:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
}
