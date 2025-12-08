package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Soal1")
	var x, y float64
	var hasil float64

	fmt.Scan(&x, &y)
	hasil = 1.0/(3.0*math.Pow(x, 2.0)+10.0) + 10.0*y + 7.0
	fmt.Println("f(x,y) = ", hasil)
}
