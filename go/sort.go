package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{425, 4, 16, 11, 43, 2}
	fmt.Println("Original:", intSlice)

	// Sorting in ascending order
	sort.Ints(intSlice)
	fmt.Println("Ascending Order:", intSlice)

	// Sorting in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println("Descending Order:", intSlice)
}
