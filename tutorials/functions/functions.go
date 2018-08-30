package functions

import (
	"fmt"
	"math"
)

//2 int parametes - type after the name
//int return type
func add(x int, y int) int {
	return x + y
}

func printAdd() {
	fmt.Println(add(42, 13))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionValues() {
	hypot := func(x, y float64) float64 { //function as a value
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot)) //function being passed around
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int { //returning a closure
	sum := 0
	return func(x int) int {
		sum += x
		return sum //Each closure is bound to its own sum variable
	}
}

func closure() { //Go functions may be closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func Run() {
	printAdd()
	functionValues()
	closure()
}
