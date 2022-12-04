package main

import "fmt"

func main() {
	//todo definisi struct
	/*Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau
	method), yang dibungkus sebagai tipe data baru dengan nama tertentu(objectstruct)
	.mirip class tapi bukan class Go menggunakan terminologi prosedural*/

	//todo panggil struct manusia
	// instansiasi object struct
	var m1 manusia
	// inisialisasi attribut object struct
	m1.nama = "Alan Turing"
	m1.umur = 41

	fmt.Println("nama :", m1.nama)
	fmt.Println("umur :", m1.umur)

	//todo inisialisasi object struct `deklarasi dgn cara berbeda`
	m2 := manusia{"steven jobs", 56} // style constructor

	var m3 = manusia{nama: "wozniak"}

	fmt.Println("manusia 2 :", m2.nama, m2.umur)
	fmt.Println("manusia 3 :", m3.nama)

	var m4 = manusia{nama: "bill gates", umur: 70}
	var m5 = manusia{umur: 50, nama: "davinci"}

	fmt.Println("manusia 4 :", m4.nama, m4.umur)
	fmt.Println("manusia 5 :", m5.nama, m5.umur)

	fmt.Println()

	//todo variabel object *pointer
	var newManusia = manusia{nama: "einstein", umur: 50}

	// pointer pManusia
	var pManusia *manusia = &newManusia
	fmt.Println("newManusia nama :", newManusia.nama, newManusia.umur) //einstein
	fmt.Println("pManusia nama :", pManusia.nama)                      //einstein

	pManusia.nama = "tesla"
	fmt.Println("newManusia nama :", newManusia.nama, newManusia.umur) //tesla
	fmt.Println("pManusia nama :", pManusia.nama)                      //tesla

	/*Meskipun pManusia bukan variabel asli, property nya tetap bisa diakses seperti biasa.
	Inilah keistimewaan property dalam objek pointer, tanpa perlu di-dereferensi nilai
	asli property tetap bisa diakses. Pengisian nilai pada property tersebut juga bisa
	langsung menggunakan nilai asli, contohnya seperti pManusia.nama = "tesla" .
	*/
	fmt.Println()

	// todo penerapan embedded struct `seperti inheritance`
	var s1 = student{}
	s1.name = "popuri" //name attribut dari person
	s1.age = 18        //age attribut dari person
	s1.grade = 2       //attribut dari student

	fmt.Println("name :", s1.name)      //popuri
	fmt.Println("age :", s1.age)        //18
	fmt.Println("age :", s1.person.age) //18
	fmt.Println("name :", s1.grade)     //2
	fmt.Println()

	//todo penerapan embedded struct dan property
	var siswa1 = siswa{}
	siswa1.name = "retsu"
	siswa1.age = 20
	siswa1.person.age = 22
	siswa1.nilai = 90

	fmt.Println(siswa1.name)       //retsu
	fmt.Println(siswa1.age)        //20
	fmt.Println(siswa1.person.age) //22
	fmt.Println(siswa1.nilai)      //90
	fmt.Println()

	//todo inisialisasi nilai sub-struct
	var p1 = person{name: "Nico Robin", age: 35}

	//Pada deklarasi z1 , property person diisi variabel objek p1 .
	var z1 = student{person: p1, grade: 4}

	fmt.Println("name : ", z1.name)  //name :  Nico Robin
	fmt.Println("age : ", z1.age)    //age :  35
	fmt.Println("grade :", z1.grade) //grade : 4
	fmt.Println()

	//todo anonymous struct
	var y1 = struct {
		person
		grade int
	}{} // anonymous struct tanpa pengisian property
	y1.person = person{"Clown", 21} //inisialisasi anonymous
	y1.grade = 2

	fmt.Println("name y1 :", y1.person.name)
	fmt.Println("age y1 :", y1.person.age)
	fmt.Println("grade  y1:", y1.grade)
	fmt.Println()

	// anonymous struct dengan pengisian property
	var element = struct {
		name      string
		particles int
	}{
		name:      "Pikachu",
		particles: 400,
	}
	fmt.Println(element)
	fmt.Println()

	//todo kombinasi slice dan struct
	var allStudents = []person{
		{name: "alien", age: 20},
		{name: "monster", age: 30},
		{name: "hero", age: 40},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
	fmt.Println()
	//todo inisialisasi slice anonymous struct
	var allstudents = []struct {
		person
		grade int
	}{
		{person: person{"mr.One", 31}, grade: 2},
		{person: person{"mr.Two", 42}, grade: 3},
		{person: person{"mr.Three", 53}, grade: 4},
	}
	for _, studentx := range allstudents {
		fmt.Println("slice anonymous struct :", studentx)
	}
	fmt.Println()

	//todo deklarasi anonymous struct dgn keyword var
	var siswa struct {
		person
		grade int
	}
	siswa.person = person{"tawon", 23}
	siswa.grade = 3
	fmt.Println("anonymous struct dgn keyword var", siswa.person, siswa.grade)

	//hanya deklarasi
	var siswax1 struct {
		grade int
	}
	// deklarasi sekaligus inisialisasi
	var siswax = struct {
		grade int
	}{
		12,
	}
	fmt.Println(siswax1, siswax)

	var p2 = struct {
		name string
		age  int
	}{name: "foo", age: 44}
	var p3 = struct {
		name string
		age  int
	}{"bar", 55}
	fmt.Println("cetak p1 dan p2 anonymous deklar :", p2, p3)
	fmt.Println()

	//todo penerapan type alias
	// pakai struct orinya atau aliasnya tetap sama
	var org = persond{"bung", 66}
	fmt.Println("cetak org persond :", org)
	var aorg = seorang{"bang", 77}
	fmt.Println("cetak aorg seorang :", aorg)

	var value number = 30
	fmt.Println("cetak value dgn typedata alias number : ", value)

}

// todo struct penerapan
type manusia struct {
	nama string
	umur int
}

// todo embedded struct
// struct person
type person struct {
	name string
	age  int
}

//struct student `extend person`
type student struct {
	grade  int
	person // embedded struct
}

// todo embedded struct dengan property yg sama
type siswa struct {
	person
	age   int // embedded property
	nilai int
}

//todo nested struct `struct yang di-embed ke sebuah struct`
/*Teknik ini biasa digunakan ketika decoding data json yang struktur datanya cukup
kompleks dengan proses decode hanya sekali.
*/
type pelajar struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

//todo deklarasi dan inisialisasi struct secara horizontal
type personx struct {
	name    string
	age     int
	hobbies []string
}

//todo tag property dalam struct
/*Tag biasa dimanfaatkan untuk keperluan encode/decode data json.*/
type persont struct {
	name string `tag1`
	age  int    `tag`
}

//todo type alias
type persond struct {
	name string
	age  int
}
type seorang = persond //pembuatan alias struct

type number = int //todo membuat alias typedata
