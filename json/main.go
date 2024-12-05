package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Definindo uma estrutura que representa os dados a serem manipulados
type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Ativo bool   `json:"ativo"`
}

func main() {
	// JSON.stringify - Convertendo uma estrutura Go para uma string JSON
	pessoa := Pessoa{
		Nome:  "João",
		Idade: 30,
		Ativo: true,
	}

	// Convertendo para JSON
	pessoaJSON, err := json.Marshal(pessoa)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pessoa em JSON (Stringify): %s\n", pessoaJSON)

	// JSON.parse - Convertendo uma string JSON de volta para uma estrutura Go
	jsonStr := `{"nome":"Maria","idade":25,"ativo":false}`

	var pessoa2 Pessoa
	// Convertendo JSON para estrutura Go
	err = json.Unmarshal([]byte(jsonStr), &pessoa2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPessoa após JSON.parse: %+v\n", pessoa2)
}
