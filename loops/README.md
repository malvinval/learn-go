## Fun fact

Dalam bahasa pemrograman Go, cuma ada 1 jenis loop yaitu `for` loop. Ga ada looping lain seperti `while`, dan `do-while`.

## for loop

Syntax:

`for var int i = 0; i < 10; i++ {}`
atau
`for i := 0; i < 10; i++ {}`

Contoh:

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

**output:**

```bash
0
1
2
3
4
5
6
7
8
9
```