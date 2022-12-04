package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari
		variabel tersebut atau bahkan memanipulasinya.

		Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang
	paling penting untuk diketahui, yaitu reflect.ValueOf() dan reflect.TypeOf() .
	`Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe
	reflect.Value , yang berisikan informasi yang berhubungan dengan nilai
	pada variabel yang dicari

	`Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe
	reflect.Type . Objek tersebut berisikan informasi yang berhubungan dengan
	tipe data variabel yang dicari
	*/

	//todo mencari tipe data dan value menggunakan reflect
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel : ", reflectValue.Type()) // int

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel : ", reflectValue.Int()) // 23
	}
	fmt.Println()

	//todo penagaksesan nilai dalam bentuk interface{}
	var angka = 10
	var reflectAngka = reflect.ValueOf(angka)

	fmt.Println("tipe variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectAngka.Interface())
	fmt.Println()

	//todo penerapan pengaksesan informasi properti variabel object
	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()
	fmt.Println()

	//todo penerapan pengaksesan informasi method variabel objek
	var s2 = &student{Name: "shigaraki tomura", Grade: 2}
	fmt.Println("nama :", s2.Name)

	var reflects2 = reflect.ValueOf(s2)
	var method = reflects2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("tomura"),
	})

	fmt.Println("nama :", s2.Name)

}

//todo pengaksesan informasi properti variabel object
type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai :", reflectValue.Field(i).Interface())
	}

}

//todo pengaksesan informasi method variabel objek
func (s *student) SetName(name string) {
	s.Name = name
}
