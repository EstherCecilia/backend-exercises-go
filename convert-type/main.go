package main

import (
	"fmt"
	"strconv"
)

func main() {
	var intVal int = 42
	floatVal := float64(intVal) // Converte int para float

	fmt.Printf("Int: %d, Float: %.2f\n", intVal, floatVal)

	strVal := strconv.Itoa(intVal) // Converte int para string

	fmt.Printf("Int: %d, String: %s\n", intVal, strVal)
}
