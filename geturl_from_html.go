package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %v", err)
	}

	// Create an io.Reader
	// io.Reader as the parameter to parse html body
	nodesTree, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse html body: %v", err)
	}

	var urls []string
	var traverseNodes func(n *html.Node)

	traverseNodes = func(n *html.Node) {
		// Process all children of a node
		// before moving to the next sibling node
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}

		// Look for <a> tag
		if n.Data == "a" {
			// Look for "href" attr
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					// To use `url.ResovleReference` method
					// we need to parse url to struct first
					parsedURL, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}

					absoluteURL := baseURL.ResolveReference(parsedURL)
					urls = append(urls, absoluteURL.String())
					break
				}
			}
		}
	}

	traverseNodes(nodesTree)
	return urls, nil
}
