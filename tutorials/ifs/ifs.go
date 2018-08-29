package ifs

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	//no (), mandatory {}
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//if with a short statement that will be executed before the condition
	//scope of v => if
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		//v can be used here
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Run() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
