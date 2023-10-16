package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) (result int) {
	result = a * b

	return
}

func main() {
	fmt.Println(sum(10, 7))
	fmt.Println(multiply(10, 7))
}
