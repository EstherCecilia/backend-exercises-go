package main

import (
	"fmt"
	"strings"
)

func main() {
	// String original
	text := "Bem-vindo ao mundo do Go!"

	// Converter para maiúsculas
	upperText := strings.ToUpper(text)
	fmt.Println("Texto em maiúsculas:", upperText)

	// Converter para minúsculas
	lowerText := strings.ToLower(text)
	fmt.Println("Texto em minúsculas:", lowerText)

	// Dividir a string em palavras
	words := strings.Fields(text)
	fmt.Println("Palavras na string:", words)

	// Substituir palavras
	replacedText := strings.Replace(text, "Go", "Golang", -1)
	fmt.Println("Texto após substituição:", replacedText)

	// Verificar se contém uma palavra
	if strings.Contains(text, "Go") {
		fmt.Println("A palavra 'Go' está presente no texto.")
	} else {
		fmt.Println("A palavra 'Go' não está presente no texto.")
	}

	// Contar ocorrências de um caractere
	count := strings.Count(text, "o")
	fmt.Printf("O caractere 'o' aparece %d vezes.\n", count)

	// Iterar sobre os caracteres da string
	fmt.Println("Iterando sobre os caracteres:")
	for index, char := range text {
		fmt.Printf("Posição %d: %c\n", index, char)
	}

	// Verificar prefixo e sufixo
	if strings.HasPrefix(text, "Bem") {
		fmt.Println("O texto começa com 'Bem'.")
	}
	if strings.HasSuffix(text, "Go!") {
		fmt.Println("O texto termina com 'Go!'.")
	}

	// Combinar elementos de um array em uma string
	joined := strings.Join(words, "-")
	fmt.Println("Palavras combinadas com '-':", joined)
}
