package main

import (
	"fmt"
)


func main() {
	var bilangan int

	fmt.Print("Masukkan sebuah bilangan bulat : ")
	fmt.Scan(&bilangan)

	if bilangan == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}