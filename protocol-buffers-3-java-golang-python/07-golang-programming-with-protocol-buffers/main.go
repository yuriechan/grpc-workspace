package main

import (
	"fmt"

	example_simple "./src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SimpleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you."
	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.GetId())
	// fmt.Println("The ID is: ", sm.Id) // NOT RECOMMENDED!!!
}
