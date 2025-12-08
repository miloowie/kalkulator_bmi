package main

import "fmt"

func main() {
	var p int
	fmt.Scan(&p)

	kapasitasDewasa := 5
	kapasitasKecil := 2

	maxDewasa := 3
	maxKecil := 5

	dewasa := p / kapasitasDewasa
	if dewasa > maxDewasa {
		dewasa = maxDewasa
	}

	sisa := p - (dewasa * kapasitasDewasa)

	kecil := sisa / kapasitasKecil
	if kecil > maxKecil {
		kecil = maxKecil
	}

	sisa = sisa - kecil*kapasitasKecil

	if sisa > 0 {
		fmt.Printf("Dewasa %d, kecil %d, dan %d tak berangkat\n", dewasa, kecil, sisa)
	} else {
		fmt.Printf("Dewasa %d, kecil %d\n", dewasa, kecil)
	}
}
