package main

import "fmt"

func cousin(number int) bool {
	divisors := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	count := 0

	if number == 1 {
		return true
	}

	for _, divisor := range divisors {
		if number%divisor == 0 {
			count++
		}
	}
	return count <= 2
}

func main() {
	const format = "Is %d cousin: %t\n"
	fmt.Printf(format, 3, cousin(3))
	fmt.Printf(format, 16, cousin(16))
	fmt.Printf(format, 53, cousin(53))
	fmt.Printf(format, 7, cousin(7))
	fmt.Printf(format, 4, cousin(4))
	fmt.Printf(format, 1, cousin(1))
	fmt.Printf(format, 2, cousin(2))
	fmt.Printf(format, 9, cousin(9))
	fmt.Printf(format, 10, cousin(10))
}
