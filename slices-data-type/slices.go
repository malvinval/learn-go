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

	// coba ganti data di slice meetingDays
	meetingDays[0] = "Rabu"

	// print data array dan slice
	fmt.Println(days)
	fmt.Println(meetingDays)

	// append()
	meetingDaysAppended := append(meetingDays, "Jumat")
	fmt.Println(meetingDaysAppended)
	fmt.Println(meetingDays)

	// append() yang menghasilkan array baru
	weekendAppended := append(weekend, "Sabit")
	weekendAppended[1] = "Minggu"
	fmt.Println(weekendAppended)
	fmt.Println(weekend)

	// make([]tipe_data, length, capacity)
	sliceAku := make([]int, 5, 10)

	sliceAku[1] = 24

	fmt.Println(sliceAku) // [0,24,0,0,0]

	// copy(dest, src)
	sliceDia := make([]int, len(sliceAku), cap(sliceAku))

	copy(sliceDia, sliceAku)

	fmt.Println(sliceDia) // [0,24,0,0,0]

	// deklarasi array dan slice itu beda!!
	iniArray := [...]int{1, 2, 3} // ini array
	iniSlice := []int{1, 2, 3}    // ini slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
