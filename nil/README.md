## Apa itu **`nil`**?

Dalam bahasa pemrograman Go, `nil` adalah nilai yang digunakan untuk merepresentasikan ketiadaan atau ketidakadaan data. Tipe data `nil` hanya bisa digunakan untuk pointer, channel, func, interface, map, atau slice.

Berikut adalah contoh penggunaannya pada pointer, slice, map, dan channel data type.

- Pointer

```go
var pointer *int

if pointer == nil {
    fmt.Println("Pointer ini nil")
}
```

- Slice

```go
var slice []int

if slice == nil {
    fmt.Println("Slice ini nil")
}
```

- Map

```go
var myMap map[string]int

if myMap == nil {
    fmt.Println("Map ini nil")
}
```

- Channel

```go
var myChannel chan int

if myChannel == nil {
    fmt.Println("Channel ini nil")
}
```