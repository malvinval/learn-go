package main

import "fmt"

type Person struct {
	name    string
	age     uint8
	married bool
}

func main() {
	// create an object struct from Person
	var person1 Person

	person1.name = "Delova"
	person1.age = 14
	person1.married = false

	fmt.Println(person1.name)

	person1.name = "Rossa"

	fmt.Println(person1.name)

	// initialize an object struct (without declaration-first)
	var person2 = Person{
		name:    "Wayland",
		age:     34,
		married: true,
	}

	fmt.Println(person2.name)
	fmt.Println(person2.age)
	fmt.Println(person2.married)
}
