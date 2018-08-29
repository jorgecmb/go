package constants

import "fmt"

const Pi = 3.14

func Run() {
	//use "const" keyword to declare constants
	//cannot use :=
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
