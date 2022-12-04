package main

import (
	"fmt"
	"runtime"
)

/*Goroutine mirip dengan thread thread, tapi sebenarnya bukan. Sebuah native
thread bisa berisikan sangat banyak goroutine. Mungkin lebih pas kalau goroutine
disebut sebagai mini thread.*/

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

//todo goroutine
/*Untuk menerapkan goroutine, proses yang akan dieksekusi sebagai goroutine
harus dibungkus kedalam sebuah fungsi. Pada saat pemanggilan fungsi tersebut,
ditambahkan keyword go didepannya, dengan itu goroutine baru akan dibuat
dengan tugas adalah menjalankan proses yang ada dalam fungsi tersebut.

Bisa dilihat di output, tulisan "halo" dan "apa kabar" bermunculan selangseling. Ini disebabkan karena statement print(5, "halo") dijalankan sebagai
goroutine baru, menjadikannya tidak saling tunggu dengan print(5, "apa
kabar") .
*/
func main() {
	//Fungsi ini digunakan untuk menentukan jumlah core atau processor yang
	//digunakan dalam eksekusi program.
	runtime.GOMAXPROCS(2)

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	//Fungsi ini akan meng-capture semua karakter sebelum user menekan tombol
	//enter, lalu menyimpannya pada variabel.
	fmt.Scanln(&input)

	//todo contoh Scanln
	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)

	// user inputs: "Monkey D Luffy"
	fmt.Println(s1) //Monkey
	fmt.Println(s2) //D
	fmt.Println(s3) //Luffy
}
