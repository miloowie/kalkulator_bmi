package main

import "fmt"

func main() {
	var k byte
	var hasil bool

	fmt.Scanf("%c", &k)

	hasil = (k >= '0' && k <= '9') || (k >= 'A' && k <= 'Z') || (k >= 'a' && k <= 'z')
	fmt.Println(hasil)
}
