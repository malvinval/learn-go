## **`fmt.Print()`**

Mencetak data tanpa newline.

## **`fmt.Println()`**

Mencetak data dengan newline.

## **`fmt.Printf(format, data)`**

Mencetak data dengan format. Contoh:

```go
    a := "Gophers"
    fmt.Printf("Data: %v, Type: %T\n", a, a)
```

`%v` untuk mengakses nilai variable a.

`%T` untuk mengakses tipe variable a.

`\n` untuk memberikan newline.

Ketiga aksi formatting diatas disebut dengan ***Formatting Verbs***. Ada banyak formatting verbs yang bisa dipake, dan berbeda-beda tiap tipe data. Untuk liat daftar verbs yang ada, bisa baca di [https://www.w3schools.com/go/go_formatting_verbs.php](https://www.w3schools.com/go/go_formatting_verbs.php)