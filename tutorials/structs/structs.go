package structs

import "fmt"

//Collection of fields
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1 = &Vertex{1, 2} // has type *Vertex (pointer)
)

func Run() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4 //access using dot

	p := &v //access using pointer
	p.Y = 1e9
	fmt.Println(v.X)
	fmt.Println(v.Y)

	fmt.Println(v1, p1, v2, v3)
}
