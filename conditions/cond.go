package main

import "fmt"

func main() {
	nums := [...]int{1, 2, 3, 4, 5, 6}

	// if array size is greater than 5, print "dang". otherwise, print "ding".
	if len(nums) > 5 {
		fmt.Println("dang")
	} else {
		fmt.Println("ding")
	}
}
