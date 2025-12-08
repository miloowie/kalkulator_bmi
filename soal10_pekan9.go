package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	siap := true

	for i := 1; i <= N; i++ {
		var a, b, c, d, e bool
		fmt.Scan(&a, &b, &c, &d, &e)

		if !(a && b && c && d && e) {
			siap = false
		}
	}

	fmt.Println(siap)
}
