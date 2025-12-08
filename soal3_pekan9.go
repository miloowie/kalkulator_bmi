package main

import "fmt"

func main() {
	var a, b, c int
	var hasil bool

	fmt.Scanln(&a, &b, &c)

	hasil = (a > 0 && b > 0 && c > 0) && (a+b > c && a+c > b && b+c > a)
	fmt.Println(hasil)
}
