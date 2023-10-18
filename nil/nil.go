package main

import "fmt"

func main() {
	type sl []int

	var n sl

	if n == nil {
		fmt.Println("Array is nil!")
	}
}
