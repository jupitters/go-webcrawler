package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {	
	cfg.concurrencyControl <- struct{}{}
	defer func() { 
		<-cfg.concurrencyControl 
		cfg.wg.Done()
		}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse rawCurrentURL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL': %v\n", err)
		return
	}

	if isFirst := cfg.addPageVisit(normalizedURL); !isFirst {
		return
	}

	fmt.Printf("Fazendo crawling em %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	proxURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error - getURLFromHTML: %v", err)
		return
	}

	for _, proxURL := range proxURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(proxURL)
	}
}