package main

import "fmt"

func main() {
	var nama string
	var umur int
	var tinggi float64
	// var lulus bool

	fmt.Scanln(&nama)
	fmt.Scanln(&umur)
	fmt.Scanln(&tinggi)

	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(tinggi)
	// fmt.Printn(lulus)
}
