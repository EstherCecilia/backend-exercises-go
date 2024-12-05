package main

import (
	"fmt"
	"net/http"
)

// Handler para a rota principal ("/")
func homePage(w http.ResponseWriter, r *http.Request) {
	// Configurando o cabeçalho e o corpo da resposta
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bem-vindo ao servidor Go!\n")
}

// Handler para uma rota de exemplo (/hello)
func helloPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Olá, este é um servidor HTTP em Go!\n")
}

func main() {
	// Definindo os manipuladores de rotas
	http.HandleFunc("/", homePage)       // Rota principal
	http.HandleFunc("/hello", helloPage) // Rota adicional

	// Iniciando o servidor
	fmt.Println("Servidor iniciado em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
