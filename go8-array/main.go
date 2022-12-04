package main

import "fmt"

func main() {

	//todo array `kumpulan data bertipe sama`
	//deklarasi variabel array dengan bertipe string semua
	var names [5]string
	names[0] = "Luffy"
	names[1] = "Zoro"
	names[2] = "Sanji"
	names[3] = "Nami"
	names[4] = "Ussop"

	fmt.Println("Cetak array names : ", names[0], names[1], names[2], names[3], names[4])

	//todo inisialisasi nilai variabel array string `fungsi len menghitung panjang array`
	fruits := [4]string{"apple", "grape", "banana", "strawberry"} // Horisontal style
	fmt.Println("Jumlah element \t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// vertikal style
	fruitx := [4]string{
		"kurma",
		"marjan",
		"kiwi",
		"olive",
	}
	fmt.Println("Jumlah element fruitx : \t", len(fruitx))
	fmt.Println("Isi semua element fruitx : \t", fruitx)

	//todo inisialisasi nilai awal array tanpa jumlah elemen
	var numbers = [...]int{3, 5, 7, 9, 11}
	fmt.Println("Data array \t", numbers)
	fmt.Println("Jumlah element \t", len(numbers))

	//todo array multidimensi
	angkaGenap := [2][3]int{{2, 4, 6}, {12, 14, 16}}
	fmt.Println("cetak array dimensi angka genap : \t", angkaGenap)

	angkaGanjil := [2][3]int{{1, 3, 5}, {11, 13, 15}}
	fmt.Println("cetak array dimensi angka ganjil : \t", angkaGanjil)

	//todo for loop array
	var animals = [4]string{"Hourse", "cow", "sheep", "chicken"}
	//todo for iterasi
	for i := 0; i < len(animals); i++ {
		fmt.Printf("Element %d : %s \n", i, animals[i])
	}

	fmt.Println()
	//vertical inisialisasi
	engineers := [4]string{
		"Informatic",
		"Electro",
		"civil",
		"machine",
	}
	//todo loop for-range `bisa _ indeksnya jika tidak digunakan`
	for indeks, engineer := range engineers {
		fmt.Printf("Engineer Element %d : %s \n", indeks, engineer)
	}

	fmt.Println()

	//todo array menggunakan keyword make
	colors := make([]string, 2)
	colors[0] = "red"
	colors[1] = "blue"

	fmt.Println(colors)

}
