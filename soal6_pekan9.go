package main

import "fmt"

func main() {
	var r1, r2, d int

	fmt.Scan(&r1, &r2, &d)

	if d > r1+r2 {
		fmt.Println("true")

	} else {
		fmt.Println("false")

	}
}
