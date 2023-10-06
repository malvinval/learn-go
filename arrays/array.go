package main

import "fmt"

func main() {
	/*
		Declare a fix-sized array
		Syntax: var nama_array = [ukuran_array] type_data {data, data, data}
	*/
	var days = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days[3], days[5])

	/*
		Declare an inferred array
		Syntax: var nama_array = [...] type_data {data, data, data}
	*/

	var students = [...]string{"Gopher1", "Gopher2", "Gopher3", "Gopher4", "Gopher5", "Gopher6", "Gopher7"}

	fmt.Println(students[0])
	fmt.Println(students[1])

	/*
		Array initialization

		- not initialized
		- partially initialized
		- fully initialized
	*/

	var numbers1 = [5]int8{}              // not initialized
	var numbers2 = [5]int8{1, 2, 3}       // partially initialized
	var numbers3 = [5]int8{1, 2, 3, 4, 5} // fully initialized

	fmt.Printf("%v\n", numbers1) // output: 0,0,0,0,0
	fmt.Printf("%v\n", numbers2) // output: 1,2,3,0,0
	fmt.Printf("%v\n", numbers3) // output: 1,2,3,4,5

	/*
		Initialized only specific element
	*/

	var foods = [5]string{1: "Pork", 4: "Candy"} // init only index 1 and 4

	fmt.Printf("Foods: %v,%v,%v\n", foods[1], foods[0], foods[4])
}
