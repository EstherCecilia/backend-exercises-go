package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma frase: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("Você digitou: %s", input)

	var value float64
	fmt.Print("Digite um número decimal: ")
	fmt.Scanf("%f", &value)
	fmt.Printf("Você digitou: %.2f\n", value)

	var num1, num2 int
	fmt.Print("Digite dois números separados por espaço: ")
	fmt.Scan(&num1, &num2)
	fmt.Printf("Número 1: %d, Número 2: %d\n", num1, num2)
}
