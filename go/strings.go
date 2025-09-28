package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "0123456789"
	fmt.Println(myString)

	myString = myString[:3] + myString[4:]
	fmt.Println(myString)

	myString = myString + "5" + "8" + "1"
	fmt.Println(myString)

	myString = strings.ReplaceAll(myString, "0", "")
	fmt.Println(myString)

	for index, char := range myString {
		fmt.Printf("%d: %s\n", index, string(char))
	}
}
