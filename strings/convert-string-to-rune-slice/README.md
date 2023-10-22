## Convert String ke Slice Rune

Dalam bahasa Go, kita bisa memecah sebuah string menjadi kumpulan rune dalam sebuah slice. Misal kita punya string `"Hello"`, maka akan diconvert jadi sebuah slice `['H','e','l','l','o']`.

Syntax: `strSlice := []rune(variableString)`

Contoh:

```go
str := "Hello"

strSlice := []rune(str)

fmt.Println(strSlice)
```

Output:

`[72 101 108 108 111]`

Jangan bingung sama angka-angka itu. Itu cuma ASCII Code (representasi angka dari tiap karakter).

Hint: Teknik `string to rune slice` ini cocok diterapin kalo kita mau sorting sebuah string.