## Apa itu Type Declaration?

Dalam bahasa Go, kita bisa membuat *alias* dari sebuah tipe data yang sudah ada. Untuk melakukannya, gunakan keyword `type`.

Syntax: `type namaAlias tipeData`. Contoh:

```go
func main() {
	type st string

	var name st = "Delova"

	fmt.Println(name)
}
```

Sebenarnya, ada banyak *alias* yang builtin. Contohnya adalah `byte` yang merupakan alias dari `uint8`.