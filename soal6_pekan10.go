package main

import "fmt"

func main() {
	var bulanini, bulankmrn float64
	fmt.Scan(&bulanini, &bulankmrn)

	if bulanini > bulankmrn {
		fmt.Println("Peningkatan sebesar", bulanini-bulankmrn)
	} else if bulanini < bulankmrn {
		fmt.Println("Penurunan sebesar", bulankmrn-bulanini)
	} else {
		fmt.Println("Tetap")
	}
}
