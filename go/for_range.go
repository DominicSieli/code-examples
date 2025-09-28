package main

import "fmt"

func main() {
	str := "0123"
	array := [4]int{0, 1, 2, 3}

	for i, v := range str {
		fmt.Printf("[%d]: %s\n", i, string(v))
	}

	fmt.Println()

	for i, v := range array {
		fmt.Printf("[%v]: %v\n", i, v)
	}
}
