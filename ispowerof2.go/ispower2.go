package main

import (
	"fmt"
)

func main() {

	fmt.Println(isPowerof2(4))
	fmt.Println(isPowerof2(16))
	fmt.Println(isPowerof2(35))
	fmt.Println(isPowerof2(66))
	fmt.Println(isPowerof2(64))
	fmt.Println(isPowerof2(127))
	fmt.Println("active bits of 7:", activeBits(7))
	fmt.Println("active bits of 8:", activeBits(8))
	fmt.Println("trailing zeroes of 6:", trailingzeros(6))
	fmt.Println("trailing zeroes of 8:", trailingzeros(8))
	fmt.Printf("Reversed bits of 1:%08b\n", reverseBits(0b00001111))
	fmt.Println(0b11100110)
}

func isPowerof2(n int) bool {
	return (n&(n-1) == 0) && n > 0
}
func activeBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func trailingzeros(n int) int {
	if n  == 0 {
		return 32
	}
	count := 0
	for (n & 1) == 0 {
		count++
		n >>= 1
	}
	return count
}

func reverseBits(n uint8) uint8 {
	var result uint8
	for i :=0; i< 8 ;i ++ {
		result = (result << 1 ) | (n & 1)
		n >>= 1
	}
	return result
}