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

## Array Initialization

Ada 3 jenis inisialisasi nilai array:

- not initialized
- partially initialized
- fully initialized

Example:

```go
var numbers1 = [5]int8{}              // not initialized
var numbers2 = [5]int8{1, 2, 3}       // partially initialized
var numbers3 = [5]int8{1, 2, 3, 4, 5} // fully initialized
```

Data dalam array yang **tidak** terinisialisasi, akan diisi dengan nilai **0** (asumsi tipe datanya adalah integer). Dengan begitu, jika array `numbers2` diprint akan menghasilkan output `[1, 2, 3, 0, 0]`, dan `numbers1` akan menghasilkan output `[0, 0, 0, 0, 0]`

## Initialized only specific element

Bahasa Go menyediakan fitur untuk programmer menginisialisasi data dalam array hanya pada index tertentu saja. Sisanya, akan diisi dengan nilai 0 (apabila bertipe integer). Contoh:

```go
var foods = [5]string{1: "Pork", 4: "Candy"} // init only index 1 and 4
fmt.Printf("Foods: %v,%v,%v\n", foods[1], foods[0], foods[4])
```

Output program diatas adalah:

```bash
Foods: Pork,,Candy
```