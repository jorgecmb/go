package functions

import "fmt"

//2 int parametes - type after the name
//int return type
func add(x int, y int) int {
	return x + y
}

func Run() {
	fmt.Println(add(42, 13))
}
