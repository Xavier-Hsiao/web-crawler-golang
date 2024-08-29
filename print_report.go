package main

import (
	"fmt"
	"sort"
)

type pageStruct struct {
	URL   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("\n=============================\n")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Printf("\n=============================\n")

	var pageSlice []pageStruct

	for url, count := range pages {
		pageSlice = append(pageSlice, pageStruct{URL: url, count: count})
	}

	sort.Slice(pageSlice, func(i, j int) bool {
		return pageSlice[i].count > pageSlice[j].count
	})

	for _, page := range pageSlice {
		fmt.Printf("Found %v internal links to %s\n", page.count, page.URL)
	}
}
