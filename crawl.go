package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse rawCurrentURL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse rawBaseURL '%s': %v\n", rawBaseURL, err)
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL': %v\n", err)
		return
	}

	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1

	fmt.Printf("Fazendo crawling em %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	proxURLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - getURLFromHTML: %v", err)
		return
	}

	for _, proxURL := range proxURLs {
		crawlPage(rawBaseURL, proxURL, pages)
	}
}