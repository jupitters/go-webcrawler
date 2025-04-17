package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)
type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

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
	if len(os.Args) < 2 {
		fmt.Println("Website não especificado.")
		fmt.Println("Uso: ./crawler <website>")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Argumentos em excesso.")
		fmt.Println("Uso: ./crawler <website>")
		return
	}

	baseURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("Erro no parametro URL: %v", err)
		return
	}

	cfg := &config{
		pages: make(map[string]int),
		baseURL: baseURL,
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 3),
		wg: &sync.WaitGroup{},
	}

	fmt.Printf("Iniciando crawler em: %s\n", baseURL.String())
	
	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL.String())
	cfg.wg.Wait()

	fmt.Println("\nResultados:")
	for normalizedUrl, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedUrl)
	}
}