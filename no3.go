package main

import "fmt"

func main() {
	var p, l, K, L int

	fmt.Scanln(&p, &l)

	K = 2*p + 2*l
	L = p * l

	fmt.Println(K, L)
}
