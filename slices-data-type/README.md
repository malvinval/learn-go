## Apa itu Slice

1. Slice merupakan ***reference*** dari sebuah Array.
2. Slice dengan Array saling terkoneksi. Apabila ada nilai Slice yang diubah, maka nilai dalam Array pun berubah. Ini terjadi karena Slice merupakan reference dari sebuah Array.

## Detail Slice

Slice butuh 3 data, yaitu pointer, length, dan capacity.

- Pointer merupakan data dalam Array yang "ditunjuk" oleh Slice.
- Length merupakan jumlah banyaknya data dalam Slice
- Capacity merupakan jumlah maksimal data yang dapat ditampung oleh Slice.

## Membuat Slice dari Array

```go
array[low:high] // membuat Slice dari array index low hingga index sebelum high (high - 1)
array[low:] // membuat Slice dari array index low hingga index terakhir di array
array[:high] // membuat Slice dari array index 0 hingga index sebelum high (high - 1)
array[:] // membuat Slice dari array index 0 hingga index terakhir di array
```

## Langsung contoh dah!

```go
// bikin array
days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// bikin slice
weekend := days[5:]   // [low:] dari index ke-5 sampai index terakhir
weekdays := days[:5]  // [:high] dari index ke-0 sampai ke 4 (5 - 1)
codingDays := days[:] // [:] dari index ke-0 sampai index terakhir
meetingDays := days[2:4] // [low:high] index ke-2 sampai index ke 3 (4 - 1)

// print data slice
fmt.Println(weekend)
fmt.Println(weekdays)
fmt.Println(codingDays)
fmt.Println(meetingDays)
```

**output:**

```bash
[Saturday Sunday]
[Monday Tuesday Wednesday Thursday Friday]
[Monday Tuesday Wednesday Thursday Friday Saturday Sunday]
[Wednesday Thursday]
```

## Length vs Capacity

Liat contoh kode diatas!

Slice **`weekdays`** itu pointernya menunjuk ke data dengan index ke-0 di Array yaitu `"Monday"` karena nilai lownya kita set ke 0, dan nilai highnya adalah 5

```
length = high - low
length = 5 - 0
length = 5

capacity = jumlah seluruh data dalam array - low
capacity = 7 - 0
capacity = 7
```

Sekarang kita coba hitung length dan capacity dari slice **`meetingDays`**.

```
length = high - low
length = 4 - 2
length = 2

capacity = jumlah seluruh data dalam array - low
capacity = 7 - 2
capacity = 5
```

Kesimpulannya, slice **`meetingDays`** memiliki pointer yang menunjuk ke data dengan index ke-2 (low), memiliki besar 2 data yaitu `[Wednesday Thursday]`, dan kapasitas maksimal 7 data.

## Mengubah data
Kalo data dalam Array diubah, maka data di Slice juga berubah. Begitu juga sebaliknya. Selalu inget bahwa **Slice adalah reference dari sebuah Array**. Contoh:

```go
// coba ganti data di slice meetingDays
meetingDays[0] = "Rabu"

// print data array dan slice
fmt.Println(days)
fmt.Println(meetingDays)
```

**output:**

```bash
[Monday Tuesday Rabu Thursday Friday Saturday Sunday]
[Rabu Thursday]
```

## Function **`len()`** & **`cap()`**

`len()` merupakan function untuk tau length (ukuran) data dalam sebuah slice.

`cap()` merupakan function untuk tau kapasitas maksimal data dalam sebuah slice.

## Function **`append()`**

`append()` merupakan function yang berguna untuk menambahkan data kedalam sebuah Slice. Namun, apabila kapasitas Slice sudah penuh maka akan dibuatkan Array baru. Contoh:

```go
meetingDaysAppended := append(meetingDays, "Jumat")
fmt.Println(meetingDaysAppended)
```

Operasi appending diatas ga akan menghasilkan array baru, karena slice `meetingDays()` masih punya sisa kapasitas.

```go
weekendAppended := append(weekend, "Sabit")
weekendAppended[1] = "Minggu"
fmt.Println(weekendAppended)
fmt.Println(weekend)
```

**output:**

```bash
[Saturday Minggu Sabit]
[Saturday Sunday]
```

Slice `weekend[1]` terlihat ga berubah jadi `"Minggu"` karena sebenernya yang diubah itu array baru, bukan Slice `weekend` asli.

## Function **`make()`**

Kita juga bisa bikin sebuah Slice secara langsung tanpa deklarasi Array sebelumnya dengan menggunakan function `make()`.

- Parameter 1: []tipe_data
- Parameter 2: length
- Parameter 3: capacity

```go
// make([]tipe_data, length, capacity)
sliceAku := make([]int, 5, 10)

sliceAku[1] = 24

fmt.Println(sliceAku)

// copy(dest, src)
sliceDia := make([]int, len(sliceAku), cap(sliceAku))

copy(sliceDia, sliceAku)

fmt.Println(sliceDia)
```

**output:**

```bash
[0 24 0 0 0]
[0 24 0 0 0]
```

## Deklarasi Array vs Slice

```go
// deklarasi array dan slice itu beda!!
iniArray := [...]int{1, 2, 3} // ini array
iniSlice := []int{1, 2, 3}    // ini slice

fmt.Println(iniArray)
fmt.Println(iniSlice)
```

Liat aja tanda [...] dan [] nya.