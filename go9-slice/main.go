package main

import "fmt"

func main() {
	// todo slice
	var fruits = []string{"apel", "pisang", "jambu", "melon"}
	fmt.Println(fruits[0])

	//todo perbedaan deklarasi slice dan array
	//slice
	var slices = []string{"foo", "bar"}
	fmt.Printf("Mencetak slices : %s \n", slices)
	//array
	var arraysA = [2]string{"banana", "melon"}
	var arraysB = [...]string{"papaya", "grape"}
	fmt.Printf("Mencetak arraysA : %s ,arraysB : %s \n", arraysA, arraysB)

	//todo teknik 2 indeks slice dan array
	// slice adalah menampung nilai/(setipedata) referensi
	colorSlice := []string{"red", "blue", "green", "yellow"}

	//mengambil element dari colorSlice
	newColorSliceA := colorSlice[0:3] //teknik dua indeks mengambil referensi
	newColorSliceAa := colorSlice[:3] //sama [0:3] hanya lebih simple
	newColorSliceB := colorSlice[2:4]
	newColorSliceC := colorSlice[2] // teknik satu indeks meng-Copy element/indeks

	fmt.Printf("Cetak colorArrayA : %s \n colorArraysB : %s \n", newColorSliceA, newColorSliceB)
	fmt.Println("Cetak colorArrayAa :", newColorSliceAa, "hasilnya sama seperti colorSliceA")
	fmt.Printf("Cetak colorArrayC :( %s )hasil dari element yg ter-copy", newColorSliceC)

	fmt.Println()
	//todo slice adalah tipedata referensi

	engineerSlice := []string{
		"informatika",
		"electro",
		"listrik",
		"robotik",
	}

	//inisialisasi referensi dari slice
	engineerA := engineerSlice[:4]
	engineerB := engineerSlice[1:4]

	fmt.Printf("Cetak engineerA : %s \nCetak engineerB : %s\n", engineerA, engineerB)

	engineerAa := engineerA[:3] //[informatika electro listrik]
	fmt.Println("engineerAa :", engineerAa)

	engineerBb := engineerB[1:3] //[listrik robotik]
	fmt.Println("engineerAb :", engineerBb)

	//todo fungsi len()
	programing := []string{"csharp", "golang", "typescript", "ruby"}

	// len `menghitung elemen slice`
	fmt.Println("fungsi len :", len(programing)) //4

	//todo fungsi cap()
	// cap `menghitung lebar or kapasitas maksimum`
	fmt.Println("fungsi cap :", cap(programing)) //4
	//[code, code, code, ----]
	capPrograming := programing[0:3]
	fmt.Println("fungsi len capPrograming :", len(capPrograming)) // 3 element
	fmt.Println("cetak capPrograming :", cap(capPrograming))      // 4 kapasitas/ruang dri slice

	//todo fungsi append() `Menambahkan element slice`

	coders := []string{"dart", "haxe", "scala", "java"}
	// append
	newCoders := append(coders, "basic") //tambah "basic" ke slice coders
	fmt.Println("cetak append xCoders : ", newCoders)

	//todo fungsi copy() `meng-Copy element slice`

	dst := make([]string, 3) //array deklarasi
	src := []string{"java", "c++", "golang", "ruby"}

	n := copy(dst, src) //copy src for inisialisasi to dst

	fmt.Println("cetak dst : ", dst) //[java c++ golang]
	fmt.Println("cetak src : ", src) //[java c++ golang ruby]

	fmt.Println("cetak n : ", n) // 3 karena cuma tiga value yg ter-copy sesuai len(dst)

	dstA := []string{"potato", "potato", "potato"}
	srcA := []string{"apple", "apple"}

	nA := copy(dstA, srcA)

	fmt.Println("cetak dstA : ", dstA) //[apple apple potato] dua value di mutasi
	fmt.Println("cetak srcA : ", srcA) //[apple apple]
	fmt.Println("cetak nA : ", nA)     //2

	//todo akses elemen slice dengan 3 indeks

	hewans := []string{"kuda", "zebra", "unta"}
	hewanA := hewans[0:2]
	hewanB := hewans[0:2:2]

	fmt.Println(hewans)      //[kuda zebra unta]
	fmt.Println(len(hewans)) //3
	fmt.Println(cap(hewans)) //3

	fmt.Println(hewanA)      //[kuda zebra]
	fmt.Println(len(hewanA)) //2
	fmt.Println(cap(hewanA)) //3

	fmt.Println(hewanB)      //[kuda zebra]
	fmt.Println(len(hewanB)) //2
	fmt.Println(cap(hewanB)) //2

}
