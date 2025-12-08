package main

import "fmt"

func main() {
	var belanja int
	var bersedia, kartu, diskon, cashback bool

	fmt.Scanln(&belanja, &bersedia)

	diskon = belanja >= 100000
	kartu = bersedia
	cashback = belanja >= 200000 && bersedia

	fmt.Println("Kartu ? :", kartu)
	fmt.Println("Diskon ? :", diskon)
	fmt.Println("Cashback ? :", cashback)

}
