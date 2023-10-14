## **`go build`**

Command `go build` akan menghasilkan sebuah executable file dari proses compile source file.

Syntax: `go build namafile.go`


## **`go mod init`**

Command `go mod init` dijalankan untuk melakukan inisialisasi module Go baru. Command ini akan menghasilkan sebuah file bernama `go.mod` yang berisi informasi tentang module yang dibuat, seperti nama module, versi Go, dan daftar dependencies.

## **`go.mod`**

File `go,mod` berisi informasi tentang module yang dibuat, seperti nama module, versi Go, dan daftar dependencies. File ini auto-generated apabila kita menjalankan command `go mod init nama_module`

## **`go mod tidy`**

Command ini melakukan sinkronisasi antara dependency yang terdaftar di file `go.mod` dengan dependency yang dibutuhkan di source file. Command ini akan menghapus dependency dari file `go.mod` apabila tidak dibutuhkan oleh source file, dan juga download dependency yang diperlukan oleh source file sekaligus memperbarui file `go.mod`.

## **`go run`**

Command ini digunakan untuk mengeksekusi kode program Go tanpa menghasilkan file executable.
Syntax: `go run namafile.go`

## **`go work`**

Secara default, go throw error kalo kita buka multiple projects dalam 1 workspace.
Itu rules workspace project go.

Contoh:

```bash
my-projects
├── hello-world
│   ├── go.mod
│   └── hello.go
└── variables
    ├── go.mod
    └── variables-etc.go
```

Kalo lu buka folder my-projects di IDE, artinya lu buka 2 project dalam 1 workspace. Go (gopls) gasuka itu!
Tapi, kalo lu mau supaya itu bisa dilakukan, pake command go work

Jalanin command berikut di directory parent, dalam contoh diatas yaitu my-projects.

1. go work init
2. go work use hello_world
3. go work use variables