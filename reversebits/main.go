package main

import (
	"fmt"
)

func main() {

	fmt.Println(ReverseBits(0b10001000))
	fmt.Println(ReverseBits(0b10000000))
	fmt.Println(ReverseBits(0b1010000))
	fmt.Println(ReverseBits(0b11000000))
	fmt.Println(R(0b11000000))
	fmt.Printf("%b\n", Swapbit(0b10010111))
	fmt.Println(PrintBits(6))
}

func ReverseBits(oct byte) byte {

	var result byte
	for i := 0; i < 8; i++ {
		result = result << 1
		result |= oct & 1
		oct >>= 1
	}
	return result
}

func Swapbit(oct byte) byte {
	return oct>>4 | oct<<4
}

func PrintBits(n int) string {
	s := ""
	for n > 0 {
		if n%2 == 1 {
			s = "1" + s
		} else {
			s = "0" + s
		}
		n /= 2
	}
	return s
}
func R(o byte) byte {
	return (o >> 7 & 1) | (o>>6&1)<<1 | (o>>5&1)<<2 | (o>>4&1)<<3 | (o>>3&1)<<4 | (o>>2&1)<<5 | (o>>1&1)<<6 | (o << 7)
}
