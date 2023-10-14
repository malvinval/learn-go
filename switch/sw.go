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
}
