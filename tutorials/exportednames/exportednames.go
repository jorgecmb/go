package exportednames

import (
	"fmt"
	"math"
)

//Only functions starting with a capital letter are exported
func Run() {
	fmt.Println(math.Pi)
}
