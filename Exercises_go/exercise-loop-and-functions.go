package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 3.0
	for i := 0; i < 10; i++ {

		z -= (z*z - x) / (2 * z)
		fmt.Printf("iteration %d, z = %v\n", i+1, z)
		if z == math.Sqrt(x) {
			fmt.Printf("Sqrt is math.Sqrt, final iteration is %d\n", i+1)
			break

		}

	}

	return z
}

func main() {
	x := 2.0
	fmt.Println("Finding square root of:", x)
	result := Sqrt(x)
	fmt.Println("Final result:", result)
	fmt.Println("math.Sqrt result:", math.Sqrt(x))
}
