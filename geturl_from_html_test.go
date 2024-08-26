package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLFromHTML(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		inputBody   string
		expected    []string
		errContains string
	}{
		{
			name:     "relative url",
			inputURL: "https://cpbl.com.tw",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>CPBL</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://cpbl.com.tw/path/one"},
		},
		{
			name:     "absolute url",
			inputURL: "https://cpbl.com.tw",
			inputBody: `
			<html>
				<body>
					<a href="https://other.com/path/one">
						<span>CPBL</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://other.com/path/one"},
		},
		{
			name:     "absolute and relative urls",
			inputURL: "https://cpbl.com.tw",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>CPBL</span>
					</a>
					<a href="https://other.com/path/one">
						<span>CPBL</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://cpbl.com.tw/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: `https://cpbl.com.tw`,
			inputBody: `
			<html>
				<body>
					<a>
						<span>CPBL</span>
					</a>
				</body>
			</html>
			`,
			expected: nil,
		},
		{
			name:     "no anchor tag",
			inputURL: `https://cpbl.com.tw`,
			inputBody: `
			<html>
				<body>
					<div>CPBL</div>
				</body>
			</html>
			`,
			expected: nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}
			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				if tc.errContains == "" {
					t.Errorf("Test %v - %s\n Failed: unexpected error: %v", i, tc.name, err)
					return
				} else if !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("Test %v - %s\n Failed: error doesn't contain expected string. Got: %v, Want: %v", i, tc.name, err, tc.errContains)
					return
				}
			} else if tc.errContains != "" {
				t.Errorf("Test %v - %s\n Failed: expected error containing: %v, but got no error", i, tc.name, tc.errContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
