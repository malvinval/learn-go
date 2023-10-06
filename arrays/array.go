package main

import "fmt"

func main() {
	/*
		Declare a fix-sized array
		Syntax: var nama_array = [ukuran_array] type_data {data, data, data}
	*/
	var days = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days[3], days[5])
}
