package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sarr := os.Args[1:]
	if len(sarr) == 0 {
		fmt.Println("enter a word")
		return
	}
	s := strings.Join(sarr, " ")
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	sbyte := []byte(s)
outerloop:
	for i, char := range sbyte {
		for _, v := range vowels {
			if char == v {
				fmt.Println(string(sbyte[i:]) + string(sbyte[:i]) + "ay")
				break outerloop
			}
		}
	}

	if !containsVowel(sbyte, vowels) {
		fmt.Println("No vowels")
	}
}

func containsVowel(s []byte, vowels []byte) bool {
	for _, char := range s {
		for _, v := range vowels {
			if char == v {
				return true
			}
		}
	}
	return false
}
