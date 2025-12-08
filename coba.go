package main

import "fmt"

func main() {
	var nilai int

	fmt.Scan(&nilai)

	if nilai <= 0 {
		fmt.Println(-nilai)
	} else {
		fmt.Println(nilai)
	}
}
