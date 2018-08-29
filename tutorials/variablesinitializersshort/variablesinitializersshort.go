package variablesinitializersshort

import "fmt"

func Run() {
	var i, j int = 1, 2
	//short assignment replacing a 'var' declaration
	//only available inside a function
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
