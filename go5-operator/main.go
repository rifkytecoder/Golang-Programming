package main

import "fmt"

func main() {

	//todo Operator Aritmatika
	//  +,-,*,/,%
	// ket: returnnya Numerik
	var tambah = 20 + 5
	kurang := 20 - 5
	bagi := 20 / 5
	modulus := 21 % 5

	fmt.Printf("hasil tambah :%d \n", tambah)
	fmt.Printf("hasil kurang :%d \n", kurang)
	fmt.Printf("hasil bagi :%d \n", bagi)
	fmt.Printf("hasil modulus :%d \n", modulus)

	panjang := 10
	lebar := 5
	luas := panjang * lebar

	fmt.Printf("hasi dari luas : %d \n", luas)

	//todo Operator Perbandingan /value
	// ==, !=, <, <=, >, >=
	// ket: return nya Boolean
	box := 10
	box2 := 5
	box3 := 10

	sama := (box == box2)
	fmt.Printf("sama dengan ? %t \n", sama)

	tidakSama := (box != box3)
	fmt.Printf("tidak sama dengan ? %t \n", tidakSama)

	//todo Operator Logika /kondisi
	// &&, ||, !
	// ket: return boolean
	var kiri = false
	kanan := true

	kiriAndKanan := kiri && kanan
	fmt.Printf("kiri && kanan \t (%t) \n", kiriAndKanan)

	kiriOrKanan := kiri || kanan
	fmt.Printf("kiri || kanan \t (%t) \n", kiriOrKanan)

	kiriReverse := !kiri
	fmt.Printf("kiri && kanan \t (%t) \n", kiriReverse)
}
