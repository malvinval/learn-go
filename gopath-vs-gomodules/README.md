## GOPATH

> Warning: Penggunaan GOPATH untuk setup project udah ga dianjurkan lagi sejak rilisnya Go versi 1.11. Untuk penggantinya, lebih disarankan untuk inisialisasi project Go menggunakan Go Modules.

GOPATH adalah variabel yang digunakan oleh Go sebagai rujukan lokasi di mana semua folder project disimpan, kecuali untuk project yang diinisialisasi menggunakan Go Modules.

**GOPATH** itu beda dengan **GOROOT**. GOROOT adalah sebuah root path yang menunjukkan lokasi Go terinstall di komputer kita. Untuk mengecek path dari GOPATH dan GOROOT, silahkan jalanin command `go env`.

GOPATH berisikan 3 buah sub-folder yakni `src`, `bin`, dan `pkg`. Kalo ga ada folder `src`, bikin aja sendiri. Biasanya, source files project Go yang disetup menggunakan GOPATH disimpan didalem folder `src` tersebut. Sebagai contoh kita ingin membuat project dengan nama `belajar`, maka harus dibuatkan sebuah folder dengan nama `belajar` dan ditempatkan dalam folder `src` (`$GOPATH/src/belajar`).

Lokasi folder yang akan dijadikan sebagai workspace bisa ditentukan sendiri. Kita bisa menggunakan alamat folder mana saja, bebas, tapi jangan gunakan path tempat di mana Go ter-install (tidak boleh sama dengan GOROOT). Yang penting adalah, path GOPATH harus terdaftar dan dikenali oleh OS kita. Contoh untuk pengguna Linux, harus daftarin path GOPATH ke dalam file `~/.bashrc`.

## Go Modules

Go modules merupakan manajemen dependensi resmi untuk Go. Modules ini diperkenalkan pertama kali di go1.11, sebelum itu pengembangan project Go dilakukan dalam GOPATH.

> Modules atau Module di sini merupakan istilah untuk project ya. Jangan bingung.

Untuk inisialisasi nama module itu bebas, tapi biasanya si namanya sama dengan nama folder projectnya. Misal, kita bikin folder project namanya `hello-world`, maka bisa dibikin nama modulenya `hello-world` juga dengan command `go mod init hello-world`.

> Btw, "nama module" dan "nama project" itu istilah yang sebenernya sama aja ya. Jangan bingung wkwk.