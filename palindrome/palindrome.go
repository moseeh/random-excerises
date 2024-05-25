package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "A man, a plan, a canal, Panama!"

	// Convert input to lowercase and remove non-alphanumeric characters
	input = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, input)

	// Check if input is a palindrome
	isPalindrome := true
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		if input[i] != input[j] {
			isPalindrome = false
			break
		}
	}

	fmt.Printf("Is \"%s\" a palindrome? %t\n", input, isPalindrome)
}
