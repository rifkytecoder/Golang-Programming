package go15library

import "fmt"

//todo akses modifier `exported` / public
func SayHello() {
	fmt.Println("Hello dari library")
}

// exported
func Sayhello(name string) {
	fmt.Println("Hello")
	introduce(name) // menyisipkan parameter name
}

//todo akses modifier `unexported` / private
func introduce(name string) {
	fmt.Println("nama saya", name)
}

// todo struct student
/*bisa disimpulkan bahwa untuk menggunakan
struct yang berada di package lain, selain nama stuct-nya harus berbentuk
exported, properti yang diakses juga harus exported juga.
*/
type Student struct {
	Name  string
	Grade int
}

// Anonymous struct
var Siswa = struct {
	Nama  string
	Grade int
}{}

//todo fungsi init
/*Selain fungsi main() , terdapat juga fungsi spesial, yaitu init() . Fungsi ini
otomatis dipanggil pertama kali ketika aplikasi di-run. Jika fungsi ini berada dalam
package main, maka dipanggil lebih dulu sebelum fungsi main() dieksekusi.*/
func init() {
	Siswa.Nama = "uchiha sasuke"
	Siswa.Grade = 888

	fmt.Println("------> go15library/library.go imported")
}
