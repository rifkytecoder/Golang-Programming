package main

import "fmt"

func main() {

	//todo perulangan for `counter loop iterasi`
	for i := 0; i < 5; i++ {
		fmt.Println("cetak i :", i)
	}

	fmt.Println()

	//todo for seperti while `loop kondisi` dengan argumen kondisi
	j := 0
	//while
	for j < 5 {
		fmt.Println("cetak j :", j)
		j++
	}

	fmt.Println()

	//todo for seperti do-while 'loop'
	contoh := 6
	for {
		// do
		fmt.Println("cetak sekali biar false")
		//while
		if contoh >= 5 {
			break
		}
		contoh++
	}

	//todo for tanpa argumen
	k := 0
	for {
		fmt.Println("cetak k:", k)
		k++
		if k == 5 {
			break
		}
	}

	fmt.Println()

	//todo for dengan keyword continue dan break
	//lanjut dan berhenti
	for n := 1; n <= 10; n++ {
		if n%2 == 1 {
			continue
		}

		if n > 8 {
			break
		}

		fmt.Println("cetak n : ", n)
	}

	fmt.Println()

	//todo for `loop bersarang`
	for a := 0; a < 5; a++ {
		for b := a; b < 5; b++ {
			fmt.Print(b, " ")
		}

		fmt.Println()
	}

	//todo for dengan label
outerLoop: //nama label bebas yg penting setelah nama pakai :
	for e := 0; e < 5; e++ {
		for f := 0; f < 5; f++ {
			if e == 3 {
				break outerLoop
			}

			fmt.Printf("Matriks [%d] [%d] \n", e, f)
		}
	}

}
