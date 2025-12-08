package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	y := x*x + 2*x + 1
	fmt.Println("Nilai y adalah:", y)
}
