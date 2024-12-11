package main

import "fmt"

func evenOrOdd(number int) bool {
	return number%2 == 0
}

func main() {
	fmt.Printf("Is even: %t\n", evenOrOdd(3))
	fmt.Printf("Is even: %t\n", evenOrOdd(2))
}
