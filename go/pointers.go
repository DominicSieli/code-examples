package main

import "fmt"

func square(n *int) int{
	return *n * *n
}

func main() {
	var value int = 0
	pointer := &value
	heapPointer := new(int)

	*pointer = 4
	*heapPointer = 8

	*pointer = square(pointer)
	*heapPointer = square(heapPointer)

	fmt.Println(*pointer)
	fmt.Println(*heapPointer)
}