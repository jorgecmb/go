package fors

import "fmt"

func Run() {
	//semi-colons separating statements. no ()
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//optional init and post statements
	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	//"while" loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//infinite loop
	for {
		break
	}
}
