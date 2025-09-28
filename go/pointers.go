package main

import "fmt"

func square(n *int) int {
	return *n * *n
}

func main() {
	var value int = 0
	stackPointer := &value
	heapPointer := new(int)

	*stackPointer = 4
	*heapPointer = 8

	*stackPointer = square(stackPointer)
	*heapPointer = square(heapPointer)

	fmt.Println(*stackPointer)
	fmt.Println(*heapPointer)
}
