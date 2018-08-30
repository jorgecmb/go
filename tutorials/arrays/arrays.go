package arrays

import "fmt"

func Run() {
	var a [2]string //declares an array of strings with 2 positions. cannot be resized
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
