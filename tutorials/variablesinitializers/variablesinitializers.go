package variablesinitializers

import "fmt"

//initializing one per variable
var i, j int = 1, 2

func Run() {
	//omitting the type. it will be inferred by the initializer value
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
