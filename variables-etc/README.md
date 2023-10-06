## Deklarasi Variabel

1. Menggunakan keyword `var`

```go
var variable_name int8
```

## Deklarasi dan inisialisasi nilai

1. Menggunakan keyword var

```go
var variable_name uint8 = 1
```

2. Menggunakan shorthand (`:=`)

Dengan menggunakan shorthand, compiler secara otomatis akan nentuin apa tipe datanya.
Anggep aja ini sama kaya keyword `auto` di C++.

```go
variable_name := 1
```

## Type Inference

Bahasa Go support *Type Inference*, sehingga programmer gaperlu nentuin tipe data secara manual ketika deklarasi variable. Contoh:

```go
var z = 10 // type inferred
var x uint8 = 5; // type deducted
```

## Multiple declaration in one line

1. Deklarasi banyak variable dengan tipe yang sama

```go
var a, b, c, d uint8 = 1,2,3,4
```

2. Deklarasi banyak variable dengan tipe yang berbeda

- Menggunakan keyword var, tanpa type.

    Kalo pake cara ini, gaboleh nentuin type.

    ```go
    var a, b = "Go", 2
    ```

- Menggunakan shorthand

    ```go
    a, b := "Go", 12
    ```


## int vs int32

- **`int`**

    The size of the int data type in Go depends on the architecture of the machine on which the program is run.
    On a 32-bit architecture, int has a size of 32 bits (4 bytes).
    On a 64-bit architecture, int has a size of 64 bits (8 bytes).
    Therefore, the size of the int data type varies.

- **`int32`**

    The int32 data type always has a size of 32 bits (4 bytes), no matter where the program is run.
    Therefore, int32 always has a consistent size.