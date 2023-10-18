package main

import (
	"fmt"
)

func sayFinished() {
	fmt.Println("Finished")
}

func sum(a int, b int) (result int) {

	// defer statement will be executed after sum() function executed.
	defer sayFinished()

	fmt.Println("Sum function is working...")

	result = a + b

	return
}

// here we implement panic() usage
func divide(a int, b int) (result int) {

	defer sayFinished()

	fmt.Println("Division function is working...")

	if a < 1 || b < 1 {
		// all statements after panic() will not be executed
		panic("Number is unacceptable!")
	}

	result = a / b

	return
}

func main() {
	num := 1

	// defer statement will be executed after main() function executed.
	if num == 1 {
		defer fmt.Println("Exit program.")
	}

	fmt.Printf("Result: %v\n", sum(10, 10))

	// trigger an error by dividing a number by zero
	divide(10, 0)
}
