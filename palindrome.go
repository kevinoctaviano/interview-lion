package main

import "fmt"

func palindromeWords(word string) bool {
	totalChar := len(word)
	firstChar := 0
	lastChar := totalChar - 1
	hasil := true

	for i := 0; i < totalChar; i++ {
		if word[firstChar] != word[lastChar] {
			hasil = false
		}
		firstChar++
		lastChar--
	}
	return hasil
}

func main() {
	fmt.Println(palindromeWords("kevin"))
	fmt.Println(palindromeWords("radar"))
	fmt.Println(palindromeWords("hello"))
	fmt.Println(palindromeWords("ollo"))
}