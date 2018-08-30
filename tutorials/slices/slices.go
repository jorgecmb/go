package slices

import (
	"fmt"
	"strings"
)

func sliceExample() {
	//fixed-length array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//dinamically-sized slice
	var s []int = primes[1:4] //Includes positions 1 to 3 (half-open range)
	fmt.Println(s)
}

func sliceReferenceArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	// A slice does not store any data, it just describes a section of an underlying array.
	fmt.Println(a, b)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	b[0] = "XXX"
	//Other slices that share the same underlying array will see those changes.
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	// creates the array and builds a slice that references it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceDefaults() {
	//default is zero for the low bound and the length of the slice for the high bound.
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[1:4])
	fmt.Println(s[:2])
	fmt.Println(s[1:])
}

func lengthAndCapacity() {
	// Length = len(s) = number of elements it contains.
	// Capacity = cap(s) = number of elements in the underlying array, counting from the first element in the slice

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	//Changing to over the capacity will trigger: "runtime error: slice bounds out of range"
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func nilSlices() {
	var s []int
	//A nil slice has a length and capacity of 0 and has no underlying array.
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlices() {
	//make function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5) //len(a) = 5, cap(a) = 5
	printSlice2("a", a)

	b := make([]int, 0, 5) //len(b) = 0, cap(b) = 5
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func slicesOfSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"}, //Slices can contain any type, including other slices.
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appending() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeExample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { //range returns 2 values: index and a copy of the element at that index.
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeExample2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { //You can skip the index or value by assigning to _. If you only want the index, drop the , value entirely.
		fmt.Printf("%d\n", value)
	}
}

func Run() {
	sliceExample()
	sliceReferenceArray()
	sliceLiterals()
	sliceDefaults()
	lengthAndCapacity()
	nilSlices()
	makeSlices()
	slicesOfSlices()
	appending()
	rangeExample()
	rangeExample2()
}
