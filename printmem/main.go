package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	for i, v := range arr {
		hex := hextodec(int(v))
		fmt.Print(hex)
		if i == 3 || i == 7 || i == 9 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
	for _, v := range arr {
		if int(v) < 32 || int(v) > 126 {
			fmt.Print(".")
		} else {
			fmt.Print(string(v))
		}
	}
	fmt.Println()

}
func hextodec(dec int) string {
	if dec == 0 {
		return "00"
	}
	s := "0123456789abcdef"

	s1 := ""

	for dec > 0 {
		s1 = string(s[dec%16]) + s1
		dec /= 16
	}
	return s1
}
