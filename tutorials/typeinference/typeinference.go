package typeinference

import "fmt"

func Run() {
	v := 0.2i //if the assigner type changes, the variable type changes
	fmt.Printf("v is of type %T\n", v)
}
