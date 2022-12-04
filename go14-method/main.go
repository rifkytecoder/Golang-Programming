package main

import (
	"fmt"
	"strings"
)

//todo method
/*Method adalah fungsi yang menempel pada type (bisa struct atau tipe data
lainnya). Method bisa diakses lewat variabel objek.
*/
// struct student
type student struct {
	name  string
	grade int
}

// method sayHello
func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

// method getName
func (s student) getName(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

//method biasa
func (s student) changeName1(name string) {
	fmt.Println("---> method biasa-> on changeName1, name change to->", name)
	s.name = name
}

//todo method pointer
func (s *student) changeName2(name string) {
	fmt.Println("--->method pointer-> on changeName2, name change to->", name)
	s.name = name
}

func main() {
	/*Cara mengakses method sama seperti pengaksesan properti berupa variabel.
	Tinggal panggil saja methodnya.*/
	//todo penerapan method
	var s1 = student{"sundar picai", 50}
	s1.sayHello()

	var name = s1.getName(2)
	fmt.Println("nama panggilan :", name)
	fmt.Println()

	//todo penerapan method pointer
	var s2 = student{"nobita", 20}

	s2.changeName1("doraemon")
	fmt.Println("s2 after changeName1 :", s2.name) //doraemon

	s2.changeName2("naruto")
	fmt.Println("s2 after changeName2 :", s2.name) //naruto
	fmt.Println()

	//method itu sendiri bisa dipanggil dari objek pointer maupun objek biasa.
	var s3 = student{"copper", 12} //`object biasa
	s3.sayHello()                  //halo copper
	var s4 = &student{"nami", 20}  //`object pointer
	s4.sayHello()                  //halo nami

}
