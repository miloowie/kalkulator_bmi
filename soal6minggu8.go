package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("Bilangan genap")
	} else {
		fmt.Println("Bilangan ganjil")
	}
}
