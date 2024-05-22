package main

import (
	"fmt"
)

func main() {

	fmt.Println(ReverseBits(0b10001000))
	fmt.Println(ReverseBits(0b10000000))
	fmt.Println(ReverseBits(0b1000000))
	fmt.Println(ReverseBits(0b10000000))
}

func ReverseBits(oct byte) byte {

	var result byte
	for i := 0; i < 8; i++ {
		result <<= 1
		result |= oct & 1
		oct >>= 1
	}
	return result
}
