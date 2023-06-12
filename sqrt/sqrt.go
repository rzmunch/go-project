package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for (z*z-x)/(2*z) != 0 { //When the condition is 0 close the loop
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1000000))
}
