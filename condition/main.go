package main

import "fmt"

func main() {
	// Exemplo de condicional simples
	age := 18
	if age >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}

	// Exemplo de condicional com múltiplas condições
	day := "segunda"
	if day == "sábado" || day == "domingo" {
		fmt.Println("Hoje é fim de semana!")
	} else if day == "sexta" {
		fmt.Println("Hoje é sexta-feira!")
	} else {
		fmt.Println("Hoje é um dia útil.")
	}

	// Exemplo de switch
	fruit := "maçã"
	switch fruit {
	case "banana":
		fmt.Println("É uma banana!")
	case "maçã":
		fmt.Println("É uma maçã!")
	default:
		fmt.Println("Não sei que fruta é essa.")
	}
}
