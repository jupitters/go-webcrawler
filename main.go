package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 399 {
		return "", errors.New("resposta falhou com status 4XX ou 5XX")
	}
	
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("non-HTML response: %s", contentType)
	}
	
	return string(body), nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Website não especificado.")
		fmt.Println("Uso: ./crawler <website>")
		return
	}
	if len(args) > 2 {
		fmt.Println("Argumentos em excesso.")
		fmt.Println("Uso: ./crawler <website>")
		return
	}

	rawBaseURL := os.Args[1]
	
	fmt.Printf("Iniciando crawler em: %s\n", rawBaseURL)
	
	pages := make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}