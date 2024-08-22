package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	a := args[0]
	b := args[1]

	aa := ""
	for _, char := range a {
		if char == '(' || char == ')' {
			continue
		} else {
			aa += string(char)
		}

	}
	fmt.Println(aa)
	if isempty(b) {
		return
	}

	matches := []string{}
	if containsPipe(aa) {
		matches = getmatch(aa)
	} else {
		counter := 1
		for _, c := range split(b) {
			if contains(c, aa) {
				fmt.Printf("%d: %s\n", counter, c)
				counter++
			}
		}
	}

	ss := split(b)
	counter := 1
	res := []string{}
	for _, s := range ss {
		for _, m := range matches {
			if contains(s, m) {
				res = append(res, s)
			}
		}
	}
	for _, s := range res {
		fmt.Printf("%d: %s\n", counter, s)
		counter++
	}
}

func split(s string) []string {
	res := []string{}

	tmp := ""
	for _, char := range s {
		if char == ' ' {
			if tmp != "" {
				res = append(res, tmp)
				tmp = ""
			}
		} else {
			tmp += string(char)
		}
	}
	if tmp != "" {
		res = append(res, tmp)
	}
	return res

}
func containsPipe(s string) bool {
	for _, char := range s {
		if char == '|' {
			return true
		}
	}
	return false
}
func getmatch(a string) []string {
	tmp := ""
	res := []string{}
	for _, char := range a {
		if char == ')' || char == '(' {
			continue
		}
		if char == '|' {
			res = append(res, tmp)
			tmp = ""
		} else {
			tmp += string(char)
		}
	}
	if tmp != "" {
		res = append(res, tmp)
	}
	return res
}

// his hi
func contains(a, b string) bool {
	size := len(b)
	for i := 0; i < len(a); i++ {
		size := i + size
		if size >= len(a) {
			size = len(a)
		}
		if b == a[i:size] {
			return true
		}

	}
	return false
}

func isempty(s string) bool {
	for _, c := range s {
		if c != ' ' {
			return false
		}
	}
	return true
}
