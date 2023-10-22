package main

import "fmt"

func main() {
	s := "Hello, world!"

	s = s[1 : len(s)-2]

	fmt.Println(s)
}
