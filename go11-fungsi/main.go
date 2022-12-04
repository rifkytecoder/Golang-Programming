package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// deklarasi variabel slice
	names := []string{"rifky", "tedy"}

	//todo panggil fungsi cetakPesan dgn mengisi parameter string dan slice string
	cetakPesan("hello nama saya : ", names)

	rand.Seed(time.Now().Unix()) //realtime dari dari os

	//todo panggil fungsi nilaiRandom
	nilaiRandom := randomWithRange(2, 10)

	// cetak nilai random dari 2 - 10
	fmt.Println("angka random : ", nilaiRandom)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	diameter := 15
	//todo panggil multi fungsi `dua variabel karena memiliki dua nilai kembali`
	// `(sesuai jumlah nilai balik fungsi).`
	area, lingkaran := kalkulasi(float64(diameter))

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", lingkaran)

	//todo panggil fungsi predefined return
	areaA, lingkaranA := calkulasi(float64(diameter))

	fmt.Printf("luas lingkaranA\t\t: %.2f \n", areaA)
	fmt.Printf("keliling lingkaranA\t: %.2f \n", lingkaranA)

}

//todo fungsi
//todo membuat fungsi cetakPesan dengan parameter string dan slice string
// `void return`
func cetakPesan(pesan string, argslice []string) {
	nameString := strings.Join(argslice, " ")
	fmt.Println(pesan, nameString)
}

//todo membuat fungsi random dengan range dan memiliki nilai kembali int
func randomWithRange(min, max int) int {
	value := rand.Int()%(max-min+1) + min
	return value
}

// todo cara membuat fungsi
// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType
// func randomWithRange(min int, max int) int
// func randomWithRange(min, max int) int

//todo fungsi dividerNumber `pembagian`
func divideNumber(m, n int) {
	// jika nilai n = 0 cetak invalid divider dan proses di hentikan
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return // bisa menghentikan proses
	}

	res := m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

//todo fungsi multiple return `banyak nilai balik`
// memiliki satu parameter diameter dan memiliki dua nilai balik float
func kalkulasi(d float64) (float64, float64) {

	// hitung keliling
	area := math.Pi * math.Pow(d/2, 2)
	// hitung lingkaran
	lingkaran := math.Pi * d
	// kembalikan dua nilai
	return area, lingkaran
}

//todo fungsi dengan predefined return value
// `konsepnya sama seperti func kalkulasi cma nilai baliknya didefinisi di awal`
func calkulasi(d float64) (area float64, lingkaran float64) {
	area = math.Pi * math.Pow(d/2, 2)
	lingkaran = math.Pi * d

	return
}
