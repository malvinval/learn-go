## Fix-sized Array

Syntax: **`var array_name = [array_size]data_type{data,data,data}`**
Contoh:

```go
var days = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

fmt.Println(days)
fmt.Println(days[0])
fmt.Println(days[1])
```

## Inferred Array

*Inferred array* artinya sebuah array yang dibuat tanpa mendefenisikan ukurannya sehingga array dapat membesar dan mengecil secara dinamis. Untuk melakukannya, gunakan operator **`...`**

Syntax: **`var array_name = [...]data_type{data,data,data}`**
Contoh:

```go
var students = [...]string{"Gopher1", "Gopher2", "Gopher3", "Gopher4", "Gopher5", "Gopher6", "Gopher7"}

fmt.Println(students)
fmt.Println(students[0])
fmt.Println(students[1])
```