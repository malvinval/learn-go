## Conditions

Pake `if else` dulu aja. Nih contohnya:

```go
nums := [...]int{1, 2, 3, 4, 5, 6}

// if array size is greater than 5, print "dang". otherwise, print "ding".
if len(nums) > 5 {
    fmt.Println("dang")
} else {
    fmt.Println("ding")
}
```

Statement `if` boleh ga make tanda kurung kalo memang ga perlu. Kecuali kalo kondisinya memang banyak, contohnya: 

```go
if len(nums) && (nums[0] == 1 || nums[0] != 2) {...}
```
Selain `if-else` ada juga keyword `else if`. Statement pada `else if` akan dieksekusi kalo statement `if` bernilai false. Dalam 1 percabangan, boleh ada banyak pengondisian `else if`.

Gaperlu dijelasin banyak-banyak lah ya tentang ini wkwk.