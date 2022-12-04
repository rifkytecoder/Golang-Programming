package main

import "fmt"

func main() {
	// todo pointer
	/*Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel
	yang berisi alamat memori suatu nilai*/

	//variabel pointer
	//`Nilai default variabel pointer adalah nil (kosong)`
	// var(
	// 	number *int
	// 	name *string
	// )

	var numberA int = 4
	// pointer *numberB terinisialisasi oleh alamat &numberA(0xc0000ac058)
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA) //4
	// &`referencing`  4 => 0xc0000ac058
	fmt.Println("numberA (memory address) :", &numberA) //0xc0000ac058

	fmt.Println()

	fmt.Println("numberB (memory address :", numberB) //0xc0000ac058
	// *`deferencing` 0xc0000ac058 => 4
	fmt.Println("numberB (value) :", *numberB) //4

	/*variabel numberA dan numberB menampung data dengan referensi alamat
	memori yang sama*/

	//todo efek perubahan nilai pointer
	var nilaiA int = 4
	var nilaiB *int = &nilaiA

	fmt.Println("nilaiA (value) :", nilaiA)    // 4
	fmt.Println("nilaiA (address) :", &nilaiA) //0xc00001a0b0

	fmt.Println("nilaiB (value) :", *nilaiB)  //4
	fmt.Println("nilaiB (address) :", nilaiB) //0xc00001a0b0

	fmt.Println()

	nilaiA = 5
	fmt.Println("nilaiA (value) :", nilaiA)    // 5
	fmt.Println("nilaiA (address) :", &nilaiA) //0xc00001a0b0

	fmt.Println("nilaiB (value) :", *nilaiB)  //5
	fmt.Println("nilaiB (address) :", nilaiB) //0xc00001a0b0

	//valuenya berubah tapi alamat memorynya tetap

	fmt.Println()

	//todo parameter pointer

	var angka int = 4
	fmt.Println("before :", angka)

	change(&angka, 10)
	fmt.Println("after :", angka)

}

// fungsi change `merubah pointer ke value`
func change(original *int, value int) {
	*original = value
}
