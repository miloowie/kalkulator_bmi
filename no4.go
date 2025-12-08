package main

import "fmt"

func main() {
	var r, L float64
	var pi float64 = 3.14

	fmt.Scanln(&r)

	L = 2 * pi * r

	fmt.Println(L)
}
