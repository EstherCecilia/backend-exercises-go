package main

import "fmt"

func main() {
	// 1. Declaração explícita de tipo
	var nome string = "João"
	var idade int = 25
	var altura float64 = 1.75

	// 2. Declaração implícita de tipo
	var cidade = "São Paulo" // Go infere que é string
	var preco = 19.99        // Go infere que é float64

	// 3. Declaração de várias variáveis ao mesmo tempo
	var (
		linguagem = "Go"
		nivel     = "Intermediário"
		ano       = 2024
	)

	// 4. Usando a declaração curta (apenas dentro de funções)
	estado := "SP" // Inferência de tipo, string
	dia := 15      // Inferência de tipo, int

	// 5. Variável sem inicialização (valor padrão do tipo)
	var valor int   // valor = 0 (valor padrão para int)
	var status bool // status = false (valor padrão para bool)

	// 6. Constantes
	const PI = 3.14159
	const Mensagem = "Bem-vindo ao Go!"

	// 7. Declaração de ponteiro (apontando para endereço de memória)
	var ptr *int
	ptr = &valor // `&` obtém o endereço de memória de `valor`

	// Exibindo os valores
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Cidade:", cidade)
	fmt.Println("Preço:", preco)
	fmt.Println("Linguagem:", linguagem, "| Nível:", nivel, "| Ano:", ano)
	fmt.Println("Estado:", estado, "| Dia:", dia)
	fmt.Println("Valor padrão de 'valor':", valor)
	fmt.Println("Valor padrão de 'status':", status)
	fmt.Println("Constante PI:", PI)
	fmt.Println("Mensagem:", Mensagem)
	fmt.Println("Endereço de memória de 'valor':", ptr)
	fmt.Println("Conteúdo do endereço apontado por 'ptr':", *ptr) // Acessa o valor do ponteiro
}
