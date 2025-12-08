package main

import "fmt"

func main() {
	var h1, m1, h2, m2, durasi, totaljam, menit int
	fmt.Scan(&h1, &m1, &h2, &m2)

	awal := h1*60 + m1
	akhir := h2*60 + m2

	if akhir < awal {
		akhir += 12 * 60
	}
	durasi = akhir - awal
	totaljam = durasi / 60
	menit = durasi % 60
	fmt.Println(totaljam, "jam", menit, "menit")
}
