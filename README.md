## Inisialiasi Project Go

1. Buat sebuah directory

    `mkdir my-project`

2. Masuk ke directory

    `cd my-project`

3. Jalankan command

    `go mod init namaModule`

## Aturan Default Workspace

Secara default, Go gamau buka multiple projects dalam 1 workspace. Misal:

```bash
my-projects
├── hello-world
│   ├── go.mod
│   └── hello.go
└── variables
    ├── go.mod
    └── variables-etc.go
```

Kalo directory my-projects dibuka di IDE, artinya 2 project dibuka dalam 1 workspace. Go (gopls) gasuka itu!

Solusinya adalah command `go work`

1. **`cd my-projects`**
2. **`go work init`**
3. **`go work use <hello_world>`**
4. **`go work use <variables>`**