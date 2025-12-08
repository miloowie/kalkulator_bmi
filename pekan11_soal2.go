package main

import "fmt"

func main() {
	var x, biayatambahan int

	fmt.Scan(&x)

	kg := x / 1000
	gr := x % 1000

	biayadasar := kg * 10000
	biayatambahan = gr

	fmt.Printf("%d kg + %d gr\n", kg, gr)

	if biayatambahan >= 500 && kg < 10 {
		tambahan := gr * 5
		total := biayadasar + tambahan
		fmt.Printf("Rp. %d + Rp. %d = Rp. %d\n", biayadasar, tambahan, total)
	} else if biayatambahan < 500 && kg < 10 {
		tambahan := gr * 15
		total := biayadasar + tambahan
		fmt.Printf("Rp. %d + Rp. %d = Rp. %d\n", biayadasar, tambahan, total)
	} else {
		fmt.Printf("Rp. %d + Rp. 0 = Rp. %d\n", biayadasar, biayadasar)
	}
}
