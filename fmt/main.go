package main

import "fmt"

func main() {
	// Imprimindo mensagens no console
	fmt.Println("Bem-vindo ao exemplo do pacote fmt!")
	fmt.Printf("O número %d é um inteiro.\n", 42)

	// Formatando valores em uma string
	name := "Esther"
	age := 30
	message := fmt.Sprintf("Meu nome é %s e eu tenho %d anos.", name, age)
	fmt.Println(message)

	// Entrada do usuário
	var nomeUsuario string
	var idadeUsuario int

	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nomeUsuario) // Lê a entrada e armazena em 'nomeUsuario'

	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idadeUsuario) // Lê a entrada e armazena em 'idadeUsuario'

	fmt.Printf("Olá, %s! Você tem %d anos.\n", nomeUsuario, idadeUsuario)

	// Demonstrando diferentes formatos de saída
	num := 42
	decimal := 3.14159
	fmt.Printf("Número inteiro: %d\n", num)
	fmt.Printf("Número decimal: %.2f\n", decimal)
	fmt.Printf("Número em notação científica: %.2e\n", decimal)

	// Alinhamento e preenchimento
	fmt.Printf("|%-10s|%10s|\n", "Esquerda", "Direita")
	fmt.Printf("|%-10d|%10d|\n", 123, 456)

	// Trabalhando com valores booleanos
	isActive := true
	fmt.Printf("O status é: %t\n", isActive)

	// Saída com %v e %+v para exibir structs
	type Pessoa struct {
		Nome  string
		Idade int
	}
	pessoa := Pessoa{Nome: "João", Idade: 25}
	fmt.Printf("Pessoa: %v\n", pessoa)
	fmt.Printf("Pessoa detalhada: %+v\n", pessoa)
}
