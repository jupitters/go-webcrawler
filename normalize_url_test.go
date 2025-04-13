package main

import (
	"reflect"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
        // add more test cases here
		{
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL	  string
		inputBody     string
		expected      []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name: "multiple relative URLs",
			inputURL:  "https://example.com",
			inputBody: `
			<html>
				<body>
					<a href="/about">About</a>
					<a href="/contact">Contact</a>
					<a href="/blog">Blog</a>
				</body>
			</html>
			`,
			expected: []string{"https://example.com/about", "https://example.com/contact", "https://example.com/blog"},
		},
		{
			name: "mixed URLs with fragments and queries",
			inputURL:  "https://example.com",
			inputBody: `
			<html>
				<body>
					<a href="/search?q=golang">Search</a>
					<a href="#section">Section</a>
					<a href="https://external.com/docs#chapter">Docs</a>
				</body>
			</html>`,
			expected: []string{"https://example.com/search?q=golang", "https://example.com/#section", "https://external.com/docs#chapter"},
		},
		{
			name: "no links",
			inputURL:  "https://example.com",
			inputBody: `
			<html>
				<body>
					<p>No links here</p>
				</body>
			</html>`,
			expected: []string{},
		},
		{
			name: "empty href attributes",
			inputURL:  "https://example.com",
			inputBody: `
			<html>
				<body>
					<a href="">Empty</a>
					<a>No href</a>
				</body>
			</html>`,
			expected: []string{"https://example.com"}, // Empty href typically resolves to base URL
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected){
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}