package main

import "fmt"

type Person struct {
	name    string
	age     uint8
	married bool
}

// create a receiver function to make a method for struct Person

func (p Person) sayHello() {
	fmt.Println("Hello, my name is", p.name)
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

	// call sayHello() method defined in struct Person
	person1.sayHello()
	person2.sayHello()
}
