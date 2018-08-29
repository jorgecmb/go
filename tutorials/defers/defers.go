package defers

import "fmt"

func stackedDefers() {
	fmt.Println("counting")

	//stacking, when the function returns, its deferred calls are executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}

	fmt.Println("done")
}

func simpleDefer() {
	//defers the execution of a function until the surrounding function returns
	defer fmt.Println("world")

	fmt.Println("hello")
}

func Run() {
	simpleDefer()
	stackedDefers()
}
