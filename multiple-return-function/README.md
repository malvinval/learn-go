## Multiple Return Function

Bahasa Go menyediakan fitur agar sebuah function bisa return lebih dari 1 nilai dengan tipe data yang sama maupun berbeda-beda. Contoh:

```go
func swap(x, y string, z int) (string, string, int) {
	return y, x, z
}

func main() {
	a, b, c := swap("hello", "world", 100)
	fmt.Println(a, b, c)
}
```

Output:

```
world hello 100
```

Dari contoh diatas, kita bisa paham bahwa function `swap()` menerima 3 parameter dan return 3 nilai yakni `y`, `x`, dan `z`. Dengan kode diatas, variable `a` akan memiliki nilai `"world"`, variable `x` memiliki nilai `"hello"`, dan `c` bernilai integer `100`.