package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	str := "00000000000000000000000000000000"
	strrune := []rune(str)
	for _, arg := range args {
		if arg[0] != '-' || len(arg) < 2 {
			fmt.Println("Invalid Option")
			return
		}
		for i, ch := range arg[1:] {
			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			if ch == 'h' && i == 0 {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			strrune[int(ch-'a')] = '1'
		}
	}
	for i := len(strrune) - 1; i >= 0; i-- {
		fmt.Print(string(strrune[i]))
		if i == 8 || i == 16 || i == 24 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
