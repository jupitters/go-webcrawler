package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 399 {
		return "", errors.New("resposta falhou com status 4XX ou 5XX")
	}
	
	if resp.Header.Get("content-type") != "text/html" {
		return "", errors.New("Content-Type not text/html")
	}
	
	return string(body), nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Website não especificado.")
		fmt.Println("Uso: ./crawler <website>")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("Argumentos em excesso.")
		fmt.Println("Uso: ./crawler <website>")
		os.Exit(1)
	}
	
	fmt.Printf("Iniciando crawler em: %s\n", args[1])
	html, err := getHTML(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(html)
}