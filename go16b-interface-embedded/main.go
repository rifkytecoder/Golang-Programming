package main

import (
	"fmt"
	"math"
)

type Ihitung2d interface {
	luas() float64
	keliling() float64
}

type Ihitung3d interface {
	volume() float64
}

//todo embedded interface
type Ihitung interface {
	Ihitung2d //interface yg ter-embedded
	Ihitung3d
}

/*Interface Ihitung2d berisikan method untuk kalkulasi luas dan keliling, sedang
Ihitung3d berisikan method untuk mencari volume bidang. Kedua interface
tersebut diturunkan di interface Ihitung , menjadikannya memiliki kemampuan
untuk menghitung luas, keliling, dan volume.
*/

//todo struct kubus
type kubus struct {
	sisi float64
}

//method volume
func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

//method luas
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

//method keliling
func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	/*Objek hasil cetakan struct kubus di atas, nantinya akan ditampung oleh objek
	cetakan interface hitung yang isinya merupakan gabungan interface hitung2d
	dan hitung3d .
	*/
	var bangunRuang Ihitung = &kubus{4}

	fmt.Println("===== kubus =====")
	fmt.Println("luas		:", bangunRuang.luas())
	fmt.Println("keliling	:", bangunRuang.keliling())
	fmt.Println("volume		:", bangunRuang.volume())

}

/*Variabel objek yang dicetak menggunakan struct
yang memiliki method pointer, jika ditampung kedalam variabel interface, harus
diambil referensi-nya terlebih dahulu.
`var bangunRuang Ihitung = &kubus{4}`
*/
