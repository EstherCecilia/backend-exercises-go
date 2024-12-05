package main

import "fmt"

func main() {
	// Loop "for" simples (de 1 a 5)
	fmt.Println("Contagem crescente:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Loop "for" com condição
	fmt.Println("Contagem regressiva:")
	count := 5
	for count > 0 {
		fmt.Println(count)
		count--
	}

	// Loop "for" infinito com break
	fmt.Println("Loop infinito controlado:")
	i := 0
	for {
		fmt.Println("Iteração:", i)
		i++
		if i >= 3 {
			break
		}
	}

	// Iterando sobre uma slice
	fmt.Println("Iterando sobre uma slice:")
	fruits := []string{"maçã", "banana", "laranja"}
	for index, fruit := range fruits {
		fmt.Printf("Fruta %d: %s\n", index+1, fruit)
	}
}
