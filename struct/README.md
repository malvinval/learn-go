## Deklarasi Struct

Dalam bahasa pemrograman Go, tipe data struct dideklarasi menggunakan keyword `type` dan `struct`.

Syntax: `type NamaStruct struct {}`. Contoh:

```go
type Person struct {
	name    string
	age     uint8
	married bool
}
```

## Deklarasi Object Struct

Dengan kode diatas, kita membuat sebuah struct bernama `Person` yang memiliki 3 properti yang berbeda tipe. Selanjutnya, kita bisa deklarasi sebuah *object struct* dari struct tersebut, lalu mengisi dan mengubah nilai properti dari *object struct* tersebut.

Syntax: `var namaObjectStruct Person`. Contoh:

```go
var person1 Person

person1.name = "Delova"
person1.age = 14
person1.married = false
```

Untuk mengakses properti sebuah struct, gunakan tanda titik (.)

Syntax: `namaObjectStruct.namaProperti`

## Inisialisasi Object Struct (Langsung)

Kalo males harus deklarasi dulu, langsung aja inisialisasi. Yang penting udah ada structnya.

Syntax: `var namaObjectStruct = Person{prop:value, prop:value}`. Contoh:

```go
var person2 = Person {
    name:    "Wayland",
    age:     34,
    married: true,
}

fmt.Println(person2.name)
fmt.Println(person2.age)
fmt.Println(person2.married)
```