package main

import "fmt"

func fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {
	fmt.Println("Fibonacci 10: ", fibonacci(10))
	fmt.Println("Fibonacci 5: ", fibonacci(5))
}
