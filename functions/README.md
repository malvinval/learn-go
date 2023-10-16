## Syntax Penulisan

Bahasa Go mempunyai syntax deklarasi function yang berbeda dengan bahasa C++.

Syntax: `func nama_function(param1 type, param2 type) return_type {}`

Contoh:

```go
func sum(a int, b int) int {
	return a + b
}
```

Dari contoh diatas, function `sum()` menerima 2 parameter bertipe integer dan akan return nilai bertipe int.

## Naked Return

Dalam bahasa pemrograman Go, kita bisa melakukan *naked return*. Artinya return statement dalam body function tidak secara eksplisit menentukan variable yang akan direturn.

Syntax: `func nama_function(param1 type, param2 type) (nama_variable_yang_ingin_direturn type) {return}`

Contoh: 

```go
func multiply(a int, b int) (result int) {
    result = a * b

    return
}

func main() {
	fmt.Println(multiply(10, 7))
}
```

Function `multiply()` akan return variable `result`. Maka outputnya adalah `70`.