## Slice (mengiris) String

Dalam bahasa pemrograman Go, kita mengenal teknik Slicing dengan menggunakan Slice. Slice tidak hanya berlaku untuk array aja, tapi bisa untuk string juga. Contoh:

```go
func main() {
	s := "Hello, world!"

	s = s[1 : len(s)-2]

	fmt.Println(s)
}
```

Output: `ello, worl`.

Dari contoh code diatas, kita bisa pahami bahwa kita membuat potongan string yang dimulai dari index ke 1 sampai index ke 11 (`len(s)-2`).

Hint: Dalam bahasa pemrograman Go, **tipe data string bersifat immutable (tidak dapat diubah)**. Misal, kita punya string `s := "Hello"` lalu kita coba hilangkan huruf `H` dengan cara `s[0] = ""`, maka akan muncul error:

```
cannot assign to s[0] (neither addressable nor a map index expression)
```

Teknik slicing diatas bisa kita gunakan sebagai solusi untuk menghilangkan suatu karakter dalam sebuah string.