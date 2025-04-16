package main

import (
	"fmt"
	"os"
)

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
}