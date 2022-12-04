package main

import "fmt"

func main() {
	//todo fungsi closure
	/*si Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
	Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi,
	atau bahkan membuat fungsi yang mengembalikan fungsi.
	*/

	//todo logic pencarian di dalam closure oleh variabel getMinMax
	getMinMax := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	numbers := []int{2, 3, 4, 3, 4, 2, 3}
	// cara panggilnya sma seperti fungsi biasa
	min, max := getMinMax(numbers)
	//%v template menampilkan segala jenis data
	fmt.Printf("data : %v\n min: %v\n max: %v\n", numbers, min, max)

	//todo IIFE `Closure jenis ini dieksekusi langsung pada saat deklarasinya`
	/*Biasa digunakan untuk membungkus proses yang hanya dilakukan sekali*/
	numberx := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	// fungsi filter number
	newNumberx := func(min int) []int {
		var r []int
		for _, e := range numberx {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("Original number :", numberx)
	fmt.Println("Filtered number :", newNumberx)

	//todo panggil fungsi findMax
	maxs := 3
	numberz := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	howMany, getNumbers := findMax(numberz, maxs)
	theNumbers := getNumbers()

	fmt.Println("numbers\t:", numberz)
	fmt.Printf("find \t: %d\n\n", maxs)
	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)

}

//todo closure sebagai nilai kembalian
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
