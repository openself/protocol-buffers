package main

import (
	"./src/simple"
	"fmt"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
}
