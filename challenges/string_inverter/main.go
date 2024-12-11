package main

import "fmt"

func invertString(str string) string {
	inverted := ""

	for i := len(str) - 1; i >= 0; i-- {
		inverted += string(str[i])
	}

	return inverted
}

func main() {
	fmt.Println(invertString("Hello World"))
}
