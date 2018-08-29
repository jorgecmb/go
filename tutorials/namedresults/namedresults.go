package namedresults

import "fmt"

//returns treated as variables defined at the top of the function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//naked return statement
	return
}

func Run() {
	fmt.Println(split(17))
}
