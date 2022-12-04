package main

import (
	"fmt"
	"strings"
)

func main() {
	//todo fungsi variadic `fungsi dengan parameter sejenis yang tak terbatas.`

	//todo panggil fungsi variadic
	avg := calculate(2, 4, 6, 8, 10, 2, 4, 6, 8)

	msg := fmt.Sprintf("Rata-rata : %.2f", avg) //mengembalikan dalam bentuk string

	fmt.Println(msg)

	/*Fungsi fmt.Sprintf() pada dasarnya sama dengan fmt.Printf() , hanya saja
	fungsi ini tidak menampilkan nilai, melainkan mengembalikan nilainya dalam
	bentuk string(parsing).*/

	//todo inisialisasi parameter fungsi variadic menggunakan data slice
	angka := []int{2, 4, 6, 8, 1, 3, 5, 7} //direkomndasikan slice styles
	avrg := calculate(angka...)            //menggunakan data slice
	mesg := fmt.Sprintf("Rata-rata : %.2f", avrg)

	fmt.Println(mesg)

	/* `sama hanya caranya yg berbeda`
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate(numbers...)
	atau
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)*/

	//todo panggil fungsi yourHobbies
	yourHobbies("rifky", "reading", "cooking", "drawing", "watch anime")

	hobby := []string{"sleeping", "eating", "adventure"} //slice styles
	yourHobbies("rifky", hobby...)                       // menggunakan data slice

}

//todo fungsi variadic `numbers adalh variadi ...`
//Parameter variadic memiliki sifat yang mirip dengan slice
func calculate(numbers ...int) float64 {
	var total int = 0

	//Cara pengaksesan tiap datanya juga sama, dengan menggunakan index.
	for _, number := range numbers {
		total += number
	}

	// `casting/parsing total int ke float64`
	avg := float64(total) / float64(len(numbers))

	return avg
}

//todo fungsi dgn parameter biasa dan variadic
func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ") // format string

	fmt.Printf("Hello my name is: %s\n", name)
	fmt.Printf("My hobbies are : %s\n", hobbiesAsString)
}
