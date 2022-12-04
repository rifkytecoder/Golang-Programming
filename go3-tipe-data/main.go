package main

import "fmt"

func main() {
	// todo tipe data
	/*
		byte
		uint 16,32,64
		int 16,32,64

		float 32,64

		Boolean true or false
	*/

	//type data numerik non-desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -212121212

	fmt.Printf("Bilangan positif : %d\n", positiveNumber)
	fmt.Printf("Bilangan negatif : %d\n", negativeNumber)

	//numerik desimal
	var decimalNumber = 3.14
	fmt.Printf("Bilangan desimal: %f \n", decimalNumber)
	fmt.Printf("Bilangan desimal: %.3f \n", decimalNumber)

	//boolean
	var isBool bool = true
	fmt.Printf("isBool? : %t \n", isBool)

	//string
	var message string = "pesan text"
	motiv := "Your Time is Limit"

	fmt.Printf("isi pesan : %s \n", message)
	fmt.Printf(`isi motivasiku "%s" \n`+"\n", motiv)

	//nilai zero value
	var (
		textString  string
		numberInt   int
		isBoole     bool
		numberFloat float32
	)

	//cetak nilai default typedata
	fmt.Printf(
		"string default: %s \n integer default: %d \n isBool defaul: %t \n float default: %f ",
		textString, numberInt, isBoole, numberFloat,
	)

}
