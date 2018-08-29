package conversions

import (
	"fmt"
	"math"
)

func Run() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) //requires an explicit conversion
	var z uint = uint(f)                          //requires an explicit conversion
	fmt.Println(x, y, z)
}
