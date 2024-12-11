package main

import (
	"fmt"
	"strings"
)

func validIndex(index int, validIndicator int, lenAlphabet int) int {
	count := lenAlphabet - index - validIndicator

	if count < 0 {
		return count * -1
	}

	return validIndicator + index

}

func cezarCipher(str string) string {
	const displacementIndicator int = 4
	alphabet := "abcdefghijklmnopqrstuvwyz"

	phrase := ""
	for i := 0; i < len(str); i++ {
		index := strings.Index(alphabet, string(str[i]))
		newIndex := validIndex(index, displacementIndicator, len(alphabet))
		if string(str[i]) == " " {
			phrase += " "

			continue
		}
		phrase += string(alphabet[newIndex])
	}

	return phrase
}

func main() {
	fmt.Println("hello world: " + cezarCipher("hello world"))
	fmt.Printf("new york times: %s\n", cezarCipher("new york times"))
}
