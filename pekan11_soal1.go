package main

import "fmt"

func main() {
	var umur int
	var kk bool

	fmt.Scan(&umur, &kk)

	if umur >= 17 && kk {
		fmt.Println("Bisa membuat KTP")
	} else if umur >= 17 && !kk || umur < 17 && kk {
		fmt.Println("belum bisa membuat KTP")
	}
}
