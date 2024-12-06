package main

import (
	"fmt"
	"sync"
	"time"
)

// Função simulando trabalho pesado
func worker(id int, wg *sync.WaitGroup, mtx *sync.Mutex, sharedCounter *int) {
	defer wg.Done() // Indica que o trabalho desta goroutine terminou
	for i := 0; i < 3; i++ {
		mtx.Lock() // Garantindo que apenas uma goroutine acesse o recurso compartilhado por vez
		*sharedCounter++
		fmt.Printf("Worker %d incrementando contador: %d\n", id, *sharedCounter)
		mtx.Unlock()
		time.Sleep(500 * time.Millisecond) // Simula trabalho
	}
}

func main() {
	var wg sync.WaitGroup // Para esperar que todas as goroutines terminem
	var mtx sync.Mutex    // Mutex para sincronizar o acesso ao recurso compartilhado
	sharedCounter := 0    // Contador compartilhado entre goroutines

	numWorkers := 5 // Número de threads (goroutines) para criar

	fmt.Println("Iniciando workers...")
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)                               // Adiciona uma tarefa no WaitGroup
		go worker(i, &wg, &mtx, &sharedCounter) // Inicia cada worker como uma goroutine
	}

	wg.Wait() // Aguarda todas as goroutines terminarem
	fmt.Println("Todos os workers concluíram!")
	fmt.Printf("Valor final do contador: %d\n", sharedCounter)
}
