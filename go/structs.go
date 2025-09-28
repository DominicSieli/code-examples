package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	dom := Person{
		name: "Dom",
		age:  44,
	}

	fmt.Printf("Name: %v\nAge: %v\n", dom.name, dom.age)
}
