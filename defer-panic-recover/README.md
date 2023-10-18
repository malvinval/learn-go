## Defer

`defer` merupakan sebuah keyword dalam bahasa Go yang berguna untuk membuat sebuah statement dieksekusi tepat setelah statement lain dieksekusi. **`defer` berlaku dalam sebuah blok function**. Contoh:

```go
func sayFinished() {
	fmt.Println("Finished")
}

func sum(a int, b int) (result int) {
	defer sayFinished()

	fmt.Println("Sum function is working...")

	result = a + b

	return
}

func main() {
	fmt.Printf("Result: %v\n", sum(10, 10))
}
```

Output:

```
Sum function is working...
Finished
Result: 20
```

Dari contoh diatas, kita bisa lihat bahwa dalam function `sum()` kita defer function `sayFinished()`. Artinya, function `sayFinished()` akan dijalankan setelah function `sum()` selesai. Walaupun statement `defer sayFinished()` diletakkan di paling atas block function `sum()`.

> **Remember!** Defer berlaku didalam sebuah block code function. Walaupun statementnya ditulis didalam block `if-else`, itu tetap berlaku dalam scope block function. Lihat contoh dibawah!

```go
func main() {
    number := 3

    if number == 3 {
        fmt.Println("halo 1")
        defer fmt.Println("halo 3")
    }

    fmt.Println("halo 2")
}
```

Output:

```
halo 1
halo 2
halo 3
```

Dari contoh diatas, bisa dipahami bahwa statement `defer fmt.Println("halo 3")` dieksekusi paling akhir (setelah function main dieksekusi). Itulah konsep `defer`.

Statement `defer` tidak dapat dihentikan oleh statement `return`. Jadi walaupun sebuah function return nilai, statement `defer` tetap akan dieksekusi setelah function tersebut selesai dieksekusi. Namun, **statement `defer` dapat dihentikan oleh statement `os.Exit(1)`**. Contoh:

```go
func main() {
    defer fmt.Println("halo")
    os.Exit(1)
    fmt.Println("selamat datang")
}
```

Dengan kode diatas, statement `defer` tidak akan dieksekusi.

## Panic

`panic()` digunakan untuk menampilkan *stack trace error* serta menghentikan flow goroutine. Semua statement yang ada setelah `panic()` tidak akan dieksekusi kecuali statement `defer()`