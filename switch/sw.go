package main

import "fmt"

func main() {
	var bin uint8 = 3

	switch bin {
	case 0:
		fmt.Println("False")
	case 1:
		fmt.Println("True")
	default:
		fmt.Println("Unknown")
	}

	var grade uint8 = 89

	switch {
	case grade > 90:
		fmt.Println("A+")
	case grade > 60:
		fmt.Println("A")
	default:
		fmt.Println("B")
	}
}
