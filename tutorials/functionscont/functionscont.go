package functionscont

import "fmt"

//omitting parameter type
func add(x, y int) int {
	return x + y
}

func Run() {
	fmt.Println(add(42, 13))
}
