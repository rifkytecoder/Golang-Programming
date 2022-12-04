package main

import "fmt"

/*Fungsi sayHello pada file partial.go bisa dikenali meski level aksesnya
adalah unexported. Hal ini karena kedua file tersebut ( main.go dan partial.go )
memiliki package yang sama.*/
func sayHallo(name string) {
	fmt.Println("Halo Bro ", name)
}
