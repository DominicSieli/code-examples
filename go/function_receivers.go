package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person *Person) PrintPerson() {
	fmt.Printf("Name: %s\nAge: %d\n", person.name, person.age)
}

func main() {
	dominic := Person{
		name: "Dominic",
		age:  45,
	}

	dominic.PrintPerson()
}
