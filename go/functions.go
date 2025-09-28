package main

import "fmt"

func square(n int) int {
	return n * n
}

func printAll(slice []string) {
	for i, v := range slice {
		fmt.Printf("[%v][%v]\n", i, v)
	}
}

func main() {
	sqr := square(5)
	names := []string{"Dom", "Ash", "Nora", "Zack"}

	fmt.Println(sqr)
	printAll(names)
}
