package main

import "fmt"

func switchStatment(caseValue int) string {
	var returnValue string

	switch caseValue {
		case 1:
			returnValue = "one"
		case 2:
			returnValue = "two"
		case 3:
			returnValue = "three"
		case 4:
			returnValue = "four"
		default:
			returnValue = "N/A"
	}

	return returnValue
}

func main() {
	for i := 0; i < 6; i++ {
		fmt.Printf("[%v]: %v\n",i, switchStatment(i))
	}
}