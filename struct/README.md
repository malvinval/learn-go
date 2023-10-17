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

## Struct Method

Kita bisa membuat sebuah method dalam sebuah struct. Cukup deklarasi sebuah *receiver function*.

Syntax: `func (p Person) namaMethod() {}`. Contoh:

```go
// create a receiver function to make a method for struct Person
func (p Person) sayHello() {
	fmt.Println("Hello, my name is", p.name)
}

// call the method
person1.sayHello()
person2.sayHello()
```

Nama variable `p` itu sebenernya bebas. Yang penting untuk dipahamin adalah, variable `p` itu gunanya untuk mengakses struct Person sesuai dengan object struct pemanggil method.