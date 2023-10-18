package main

import "fmt"

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

func main() {
	num := 1

	// defer statement will be executed after main() function executed.
	if num == 1 {
		defer fmt.Println("Exit program.")
	}

	fmt.Printf("Result: %v\n", sum(10, 10))
}
