package main

import "fmt"

func main() {
	var a, b, c int
	var hasil bool

	fmt.Scan(&a, &b, &c)

	hasil = (a > b && a < c) || (a > c && a < b)
	fmt.Println(hasil)
}
