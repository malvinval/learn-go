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

## int vs int32

- **`int`**

    The size of the int data type in Go depends on the architecture of the machine on which the program is run.
    On a 32-bit architecture, int has a size of 32 bits (4 bytes).
    On a 64-bit architecture, int has a size of 64 bits (8 bytes).
    Therefore, the size of the int data type varies.

- **`int32`**

    The int32 data type always has a size of 32 bits (4 bytes), no matter where the program is run.
    Therefore, int32 always has a consistent size.