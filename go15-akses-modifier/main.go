package main

import (
	"fmt"
	f "fmt" // `menggunakan alias f`
	"golang/go15library"
	. "golang/go15library" // import dgn prefix (.)
)

/*Cara pemanggilan fungsi yang berada dalam package lain adalah dengan
menuliskan nama package target diikut dengan nama fungsi menggunakan dot
notation atau tanda titik, seperti go15library.SayHello()*/

func main() {
	//todo akses modifier penerapan
	go15library.SayHello() //Hello dari library
	//go15library.introduce("hoi") error karena unexported / private

	go15library.Sayhello("Golang")
	fmt.Println()

	//todo akses struct dari package lain
	/*bisa disimpulkan bahwa untuk menggunakan
	  struct yang berada di package lain, selain nama stuct-nya harus berbentuk
	  exported, properti yang diakses juga harus exported juga.
	*/
	var s1 = go15library.Student{"kaiba", 20} // tanpa menggunakan prefix

	fmt.Println("Name :", s1.Name)   // Name : kaiba
	fmt.Println("Grade :", s1.Grade) // Grade : 20

	//todo penerapan prefix
	var prefixs1 = Student{"prefix", 99} // menggunakan prefix
	fmt.Println("Nama :", prefixs1.Name)
	fmt.Println("Grade :", prefixs1.Grade)
	/*Pada kode di atas package go15library di-import menggunakan tanda titik. Dengan
	itu, pemanggilan struct Student tidak perlu dengan menuliskan nama package
	nya.*/
	f.Println()

	//todo alias f / fmt
	f.Println("Hello Alias Golang")
	f.Println()

	//todo panggil fungsi dari package yg sama `partial.go`
	/*Sekarang terdapat 2 file berbeda ( main.go dan partial.go ) dengan package
	adalah sama, main . Pada saat go build atau go run , semua file dengan
	nama package main harus dituliskan sebagai argumen command*/
	sayHallo("Yugioh") //Halo Bro  Yugioh `dari partial.go`
	//? go run main.go partial.go atau bisa go run *.go

	f.Println()
	//todo penerapan fungsi init
	/*Package library di-import, dan variabel Student dikonsumsi. Pada saat import
	package, fungsi init() yang berada didalamnya langsung dieksekusi.
	Property variabel objek Student akan diisi dan sebuah pesan ditampilkan ke
	console.
	Dalam sebuah package diperbolehkan ada banyak fungsi init() (urutan
	eksekusinya adalah sesuai file mana yg terlebih dahulu digunakan). Fungsi ini
	dipanggil sebelum fungsi main() , pada saat eksekusi program.*/
	f.Printf("Nama : %s\n", go15library.Siswa.Nama) //Nama : uchiha sasuke
	f.Printf("Grade : %d\n", Siswa.Grade)           //888

}
