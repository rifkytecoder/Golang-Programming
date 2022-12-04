package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"anna", "karen", "popuri"}
	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	dataLenght5 := filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data) //[anna karen popuri]

	fmt.Println("filter ada huruf \"o\"\t :", dataContainsO) //[popuri]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5) //[karen]

}

//todo fungsi sebagai parameter
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

/*1. Data array (yang didapat dari parameter pertama) akan di-looping.
2. Di tiap perulangannya, closure callback dipanggil, dengan disisipkan data
tiap elemen perulangan sebagai parameter.
3. Closure callback berisikan kondisi filtering, dengan hasil bertipe bool
yang kemudian dijadikan nilai balik dikembalikan.
4. Di dalam fungsi filter() sendiri, ada proses seleksi kondisi (yang nilainya
didapat dari hasil eksekusi closure callback ). Ketika kondisinya bernilai
true , maka data elemen yang sedang diulang dinyatakan lolos proses
filtering.
5. Data yang lolos ditampung variabel result . Variabel tersebut dijadikan
sebagai nilai balik fungsi filter() .
*/

/*Pada fungsi yang skema-nya cukup panjang, akan lebih baik jika menggunakan
alias, apalagi ketika ada parameter fungsi lain yang juga menggunakan skema
yang sama. Membuat alias fungsi berarti menjadikan skema fungsi tersebut
menjadi tipe data baru. Caranya dengan menggunakan keyword type . Contoh:

// type FilterCallback func(string) bool
   func filter(data []string, callback FilterCallback) []string {
 // ...
}
Skema func(string) bool diubah menjadi tipe dengan nama FilterCallback .
Tipe tersebut kemudian digunakan sebagai tipe data parameter callback .*/
