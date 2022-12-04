package main

import "fmt"

func main() {

	// todo deklari dengan type data
	//var firstName string = "rifky"

	var (
		firstName string = "rifky"
		lastName  string
	)
	lastName = "Te"

	// \n literal
	fmt.Printf("halo nama saya %s %s \n", firstName, lastName)
	//concatenation +
	fmt.Println("haloo", firstName, lastName+"!")

	fmt.Println("=================================")

	// todo deklarasi tanpa type data
	// style type inference menyesuaikan dengan value
	namaDepan := "foo"

	namaBelakang := "bar"
	namaBelakang = "new bar"

	// bisa juga
	var namaTengah = "food"

	fmt.Printf("Hay nama saya %s %s %s !", namaDepan, namaTengah, namaBelakang)

	// todo deklarasi Multi variabel
	var first, second, third string

	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	var fourt, fifth, sixth string = "empat", "lima", "enam"
	fmt.Println(fourt, fifth, sixth)

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Println(seventh, eight, ninth)

	one, isFriday, pi, hello := 1, true, 3.14, "hello world"
	fmt.Println(one, isFriday, pi, hello)

	// todo variabel underscore
	_ = "belajar Golang"
	_ = "Golang it easy"
	name, _ := "foo", "bar"
	fmt.Println(name)

	// todo deklarasi variabel mengunakan keyword new untuk membuat pointer
	human := new(string)
	fmt.Println("cetak human :", human)   //0xc000088350 pointer string/alamat memory
	fmt.Println("cetak *human :", *human) //"" gunakan * untuk melihat nilai asli variabel

}
