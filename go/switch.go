package main

import "fmt"

const (
	ZERO int = iota
	ONE
	TWO
	THREE
	FOUR
)

func switchStatment(caseValue int) string {
	var returnValue string

	switch caseValue {
	case ZERO:
		returnValue = "Zero"
	case ONE:
		returnValue = "One"
	case TWO:
		returnValue = "Two"
	case THREE:
		returnValue = "Three"
	case FOUR:
		returnValue = "Four"
	default:
		returnValue = "N/A"
	}

	return returnValue
}

func main() {
	for i := 0; i < 6; i++ {
		fmt.Printf("[%v]: %v\n", i, switchStatment(i))
	}
}
