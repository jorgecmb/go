package multipleresults

import "fmt"

//returning multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func Run() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
