package variables

import "fmt"

//var at package level declaring variables
//type is last
var c, python, java bool

func Run() {
	//var at function level
	var i int
	fmt.Println(i, c, python, java)
}
