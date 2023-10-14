package main

import "fmt"

func main() {
	// bikin array
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// bikin slice

	weekend := days[5:]      // [low:] dari index ke-5 sampai index terakhir
	weekdays := days[:5]     // [:high] dari index ke-0 sampai ke 4 (5 - 1)
	codingDays := days[:]    // [:] dari index ke-0 sampain index terakhir
	meetingDays := days[2:4] // [low:high] index ke-2 sampai index ke 3 (4 - 1)

	// print data slice
	fmt.Println(weekend)
	fmt.Println(weekdays)
	fmt.Println(codingDays)
	fmt.Println(meetingDays)
}
