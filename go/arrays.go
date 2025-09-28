package main

import "fmt"

func main() {
	array1 := [2]int{}
	array2 := new([2]int)
	var array3 [2]int = [2]int{0, 0}

	fmt.Println(array1)
	fmt.Println(*array2)
	fmt.Println(array3)
}
