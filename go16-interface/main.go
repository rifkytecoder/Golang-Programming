package main

import (
	"fmt"
	"math"
)

/*Interface(`type data`) adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi
saja), yang dibungkus dengan nama tertentu.
*/

//todo interface hitung
type Ihitung interface {
	luas() float64     //method `deklarasi method kosong`
	keliling() float64 //method
}

//todo struct lingkaran
type lingkaran struct {
	diameter float64
}

//method jariJari `dari struct lingkaran tidak terdefinisi di interface
func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

//method luas `dari struct lingkaran terdefinisi di interface
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

//method keliling `dari struct lingkaran terdefinisi di interface
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

//todo struct persegi
type persegi struct {
	sisi float64
}

//method luas `dari struct persegi terdefinisi di interface
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

//method keliling `dari struct persegi terdefinisi di interface
func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	//todo penerapan deklarasi interface
	var bangunDatar Ihitung //object interface

	//inisialisasi variabel bangunDatar
	bangunDatar = persegi{10}
	fmt.Println("==== persegi====")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())

	bangunDatar = lingkaran{14}
	fmt.Println("==== lingkaran====")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())
	fmt.Println("jari-jari	:", bangunDatar.(lingkaran).jariJari())
	/*Untuk mengakses method yang tidak ter-definisi di interface, variabel-nya harus
	di-casting terlebih dahulu ke tipe asli variabel konkritnya (pada kasus ini tipenya
	lingkaran ), setelahnya method(`jariJari`) akan bisa diakses.

	variabel interface hanya bisa
	menampung objek yang minimal memiliki semua method yang terdefinisi di
	interface-nya.*/

	//method luas dan keliling terdefinisi di interface jdi bsa langsung di panggil

}
