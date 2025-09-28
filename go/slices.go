package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	slice = append(slice[:3], slice[4:]...)
	fmt.Println(slice)

	slice = append(slice[:len(slice)-1])
	fmt.Println(slice)

	slice = append(slice, 3)
	fmt.Println(slice)

	slice = append(slice, []int{100, 200, 300}...)
	fmt.Println(slice)
}
