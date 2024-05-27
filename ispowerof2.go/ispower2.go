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
}

func isPowerof2(n int) bool {
	return (n&(n-1) == 0) && n > 0
}
