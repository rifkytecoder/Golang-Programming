package main

import "fmt"

func main() {
	//todo map tipe data data key-value

	//deklarasi tipe map
	cobaMap := map[string]int{} //[key]value

	//assign value ke map
	cobaMap["januari"] = 10
	cobaMap["februari"] = 30

	fmt.Println("key Januari : ", cobaMap["januari"]) //key Januari :  10
	fmt.Println("april :", cobaMap["april"])          //april : 0

	//var data map[string]int //nil
	// data["one"] = 1
	// fmt.Println(data["one"])
	data := map[string]int{} //inisial explisit
	data["one"] = 1

	fmt.Println("Cetak data map : ", data["one"])

	// style inisialisasi vertikal
	var dataA = map[string]int{"apple": 10, "pear": 20, "cerry": 30}

	// style inisialisasi horizontal
	var dataB = map[string]int{
		"java":   11,
		"python": 21,
		"ruby":   31,
		"golang": 42,
	}
	fmt.Printf("dataA : %d \n", dataA["apple"]) //cetak per key
	fmt.Printf("dataB : %d \n", dataB["golang"])
	fmt.Println(dataA) //cetak semua map
	fmt.Println(dataB)

	//todo gaya pembuatan map yg sama
	//	var Amap = map[string]int{}
	//	var Bmap = make(map[string]int)
	//	var Cmap = *new(map[string]int)

	months := map[string]int{
		"januari":  100,
		"februari": 200,
		"maret":    300,
		"april":    400,
	}
	//iterasi for
	for key, val := range months {
		fmt.Printf("cetak map months : key (%s) value:(%d) \n", key, val)
	}
	fmt.Println(len(months)) //mencetak panjang dari map

	//todo fungsi delete hapus
	delete(months, "januari") // key januari akan dihapus
	fmt.Println("cetak map months yg sudah di delete : \n", months)

	//todo deteksi item dengan item tertentu
	values := map[string]int{
		"satu":  1,
		"dua":   2,
		"tiga":  3,
		"empat": 4,
		"lima":  5,
	}

	//var ke-2 nilai bool
	var val, isExist = values["enam"] //false condition
	if isExist {
		fmt.Println(val)
	} else {
		fmt.Println("item is not exists")
	}

	//todo kombinasi slice dan map

	//slice dan map kombinasi
	fruits := []map[string]string{
		{"name": "apple", "warna": "merah"},
		{"name": "banana", "warna": "kuning"},
		{"name": "grape", "warna": "ungu"},
	}
	for _, fruitx := range fruits {
		//cetak semua value
		fmt.Println(fruitx["name"], ":", fruitx["warna"])
	}

	//slice map kombinsasi dengan key-value random
	datax := []map[string]string{
		{"nama": "sapi", "gendre": "betina", "warna": "coklat"},
		{"alamat": "manado", "id": "k001"},
		{"komunitas": "sapi lovers"},
	}

	//cetak semua key dan value
	for _, dataxs := range datax {
		fmt.Println(dataxs)
	}

}
