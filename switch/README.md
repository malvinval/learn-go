## Switch

Bisa dibilang ini saudara dari pengondisian `if else`. Kita langsung bikin contoh aja deh.

```go
var bin uint8 = 3

switch bin {
case 0:
    fmt.Println("False")
case 1:
    fmt.Println("True")
default:
    fmt.Println("Unknown")
}
```

Dari kode diatas, kita akan membuat pengondisian dari nilai variabel `bin`. Kalo nilainya 0, akan diprint `"False"`. Kalo nilainya 1, akan diprint `"True"`. Diluar itu, maka akan diprint `"Unknown"`.

Selain itu juga, percabangan dengan switch case bisa juga digunakan layaknya seperti memakai percabangan if else, caranya adalah sebagai berikut:

```go
var grade uint8 = 89

switch {
case grade > 90:
    fmt.Println("A+")
case grade > 60:
    fmt.Println("A")
default:
    fmt.Println("B")
}
```

**Contoh pertama syntaxnya adalah `switch(variabel){}`, sedangkan yang kedua adalah `switch {}`**