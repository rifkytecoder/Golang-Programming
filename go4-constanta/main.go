package main

import "fmt"

func main() {
	const firstName string = "Rifky"
	fmt.Printf("My name is %s \n", firstName)

	const lastName = "Te"
	fmt.Printf("Nama belakang saya %s \n", lastName)

	fmt.Print("foo", "bar \n")   //spasi manual
	fmt.Println("food", "water") //spasi automatis
}
