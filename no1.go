package main

import "fmt"

func main() {
	var nilai int
	var tugas bool

	fmt.Scan(&nilai, &tugas)

	lulus := (nilai > 55 && tugas) || (nilai > 90)
	print(lulus)
}
