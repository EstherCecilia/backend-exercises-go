package main

import "fmt"

func ordination(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				aux := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = aux
			}
		}
	}
	return numbers
}

func main() {
	fmt.Println(ordination([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(ordination([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(ordination([]int{1, 3, 5, 7, 9, 2, 4, 6, 8}))
	fmt.Println(ordination([]int{9, 7, 5, 3, 1, 8, 6, 4, 2}))
}
