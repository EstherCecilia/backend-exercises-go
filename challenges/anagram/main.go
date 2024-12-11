package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(srt string, anagramSrt string) bool {
	if len(srt) != len(anagramSrt) {
		return false
	}

	arraySrt := strings.Split(srt, "")
	arrayAnagramSrt := strings.Split(anagramSrt, "")

	sort.Strings(arraySrt)
	sort.Strings(arrayAnagramSrt)

	if strings.Join(arraySrt, "") == strings.Join(arrayAnagramSrt, "") {
		return true
	}

	return false
}

func main() {
	fmt.Println(anagram("hello", "olleh"))
	fmt.Println(anagram("hello", "world"))
	fmt.Println(anagram("hello", "olleh"))
	fmt.Println(anagram("hello", "olehl"))
	fmt.Println(anagram("hello", "hello"))
}
