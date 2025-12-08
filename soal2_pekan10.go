package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("kelipatan 3", "kelipatan 5")
	} else if x%3 == 0 {
		fmt.Println("kelipatan 3")
	} else if x%5 == 0 {
		fmt.Println("kelipatan 5")
	} else {
		fmt.Println()
	}
}
