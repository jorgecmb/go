package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	lastZ := 0.0
	count := 1

	for z-lastZ > 0 {
		lastZ = z
		z -= (z*z - x) / (2 * z)
		count++
	}

	fmt.Println(count)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
