package main

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.New("não foi possivel transformar em URL")
	}

	formatedURL := parsedURL.Host + strings.TrimRight(parsedURL.Path, "/")
	
	return formatedURL, nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	URLs := []string{}

	parsedHttp, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, errors.New("não foi possivel ler o HTML")
	}
	// absoluteURL, err := normalizeURL(rawBaseURL)
	// if err != nil {
	// 	return []string{}, errors.New("não foi possivel obter a URL absoluta")
	// }

	for n := range parsedHttp.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if strings.Index(a.Val, "/") == 0{
						a.Val = rawBaseURL + a.Val
					}
					if strings.Index(a.Val, "#") == 0 {
						a.Val = rawBaseURL + "/" + a.Val
					}
					URLs = append(URLs, a.Val)
					//fmt.Println(a.Val)
					break
				}
			}
		}
	}

	return URLs, nil
}