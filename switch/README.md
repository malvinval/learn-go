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