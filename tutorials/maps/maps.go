package maps

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

func mapExample() {
	var m map[string]Vertex
	m = make(map[string]Vertex) //make initializes the map
	m["Bell Labs"] = Vertex{    //adding a key-value
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	var m = map[string]Vertex{ //like structs with keys
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mapLiterals2() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967}, //you can omit the type name on elements
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mutate() {
	m := make(map[string]int)

	m["Answer"] = 42 //Insert or update an element in map
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") //Delete an element
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] //Test that a key is present with a two-value assignment
	fmt.Println("The value:", v, "Present?", ok)
}

func wordCount(s string) map[string]int {
	m := make(map[string]int)   //building the map
	fields := strings.Fields(s) //splitting the string into words

	for _, v := range fields {
		value := m[v]
		m[v] = value + 1
	}
	return m
}

func Run() {
	mapExample()
	mapLiterals()
	mapLiterals2()
	mutate()
	fmt.Println(wordCount("This is a string with some words but it is still a string"))
}
