package main

import (
	"fmt"
)

/*
	- Variables in go is static typed.
	- Declaration format: var variable_name variable_type = value;
	- Basic data types in go:

	bool: true / false

	string: collection of chars

	int
	int8: range -128 through 127.
	int16: range -32768 through 32767
	int32: range -2147483648 through 2147483647
	int64: range -9223372036854775808 through 9223372036854775807.

	uint
	uint8: range 0 through 255 (128+127)
	uint16: range 0 through 65535 (32768+32767)
	uint32: range 0 through 4294967295 (2147483648 + 2147483647)
	uint64: range 0 through 18446744073709551615 (9223372036854775808 + 9223372036854775807)

	uintptr

	byte:
		- alias for uint8

	rune:
		- alias for int32
		- represents a Unicode code point

	float32
	float64

	complex64
	complex128
*/

// Declare multiple variables simultaneously
var (
	age        uint8  = 19
	name       string = "Delova Rossenberg"
	bloodTypeR rune   = 'A'

	// convert rune to string
	bloodTypeS string = string(bloodTypeR)
)

func main() {
	// using shorthand
	height := 182.4

	// Multiple declaration in one line
	var a, b, c, d int = 1, 3, 5, 7

	// If the type keyword is not specified, you can declare different types of variables in the same line
	var e, f = 6, "Hello"
	g, h := 7, "World!"

	fmt.Println(name, "is", age, "years old with", bloodTypeS, "bloodtype.")
	fmt.Printf("Height: %v centimeters\n", height)
	fmt.Println(a, b, c, d, e, f, g, h)

}
