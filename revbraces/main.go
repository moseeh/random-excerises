package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "foo(bar)bazblim"
	s = reverseInnermost(s)
	s = strings.NewReplacer("(", "", ")", "").Replace(s)
	fmt.Println(s)
}

func reverseInnermost(s string) string {
	srune := []rune(s)
	stack := []int{}

	for i, v := range srune {
		if v == '(' {
			stack = append(stack, i)
		} else if v == ')' && len(stack) > 0 {
			start := stack[len(stack)-1] + 1
			end := i
			reverse(srune[start:end])
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}
		}
	}

	return string(srune)
}

func reverse(r []rune) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}
