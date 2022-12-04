package main

import (
	"fmt"
	"strings"
)

func main() {
	/*Interface kosong atau empty interface yang dinotasikan dengan interface{}
	merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung
	segala jenis data, bahkan array, pointer, apapun. Tipe data dengan konsep ini
	biasa disebut dengan dynamic typing.*/
	//todo interface kosong{} atau type data any
	var secret interface{} // jdi type data

	secret = "kuroko"
	fmt.Println(secret)

	secret = []string{"mangga", "pisang", "manggis"}
	fmt.Println(secret)

	secret = 3.14
	fmt.Println(secret)
	fmt.Println()

	//todo map string : interface{}
	var data map[string]interface{}

	//inisialisasi value dengan berbagai jenis type data
	data = map[string]interface{}{
		"name":      "rafaela",                             //string
		"grade":     2,                                     //int
		"breakfast": []string{"apple", "manggo", "banana"}, //slice
	}

	// cetak dgn iterasi for range
	for _, vdata := range data {
		fmt.Println(vdata)
	}
	fmt.Println()

	//todo casting variabel interface kosong
	/*Hal ini penting diketahui, karena untuk melakukan operasi yang membutuhkan
	nilai asli pada variabel yang bertipe interface{} , diperlukan casting ke tipe
	aslinya.*/
	//Teknik casting pada interface disebut dengan type assertions.
	var rahasia interface{}

	rahasia = 2
	var number = rahasia.(int) * 10 // casting / parsing ke int
	fmt.Println(rahasia, "multiplied by 10 is :", number)

	rahasia = []string{"Jeruk", "Lemon", "Pear"}
	var fruits = strings.Join(rahasia.([]string), ", ") // parsing ke string
	fmt.Println(fruits, "-Buah favorit ku ")
	fmt.Println()

	//todo casting variabel interface kosong ke object pointer
	//Variabel interface{} bisa menyimpan data apa saja, termasuk data objek,pointer, ataupun gabungan keduanya.
	var xsecret interface{} = &person{name: "wick", age: 27}
	var name = xsecret.(*person).name //casting ke struct pointer
	fmt.Println(name)
	fmt.Println()

	//todo kombinasi slice, map, dan interface{}
	/* variabel person dideklarasikan berisi data slice map
	berisikan 2 item dengan key adalah name dan age .
	*/
	var person = []map[string]interface{}{
		{"nama": "Go", "age": 11},
		{"nama": "Typescript", "age": 10},
		{"nama": "Kotlin", "age": 11},
	}
	for _, vperson := range person {
		fmt.Println(vperson["nama"], "age is", vperson["age"])
	}
	fmt.Println()

	/*Dengan memanfaatkan slice dan interface{} , kita bisa membuat data array
	yang isinya adalah bisa apa saja.*/
	var buah = []interface{}{
		map[string]interface{}{"name": "Pisang", "total": 10},
		[]string{"Manggis", "Semangka", "Duku"},
		"Orange",
	}
	for _, vbuah := range buah {
		fmt.Println(vbuah)
	}
}

//`alias`type Any = interface{}

type person struct {
	name string
	age  int
}
