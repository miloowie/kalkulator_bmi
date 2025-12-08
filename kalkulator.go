package main

import "fmt"

func main() {
	var bb, tb float64
	var gender string

	fmt.Println("Kalkulator BMI")

	for {
		fmt.Print("Jenis Kelamin (L/P): ")
		fmt.Scan(&gender)

		if gender == "L" || gender == "P" {
			break
		}
		fmt.Println("Input tidak valid!")
	}

	for {
		fmt.Print("Berat Badan (kg): ")
		_, err := fmt.Scan(&bb)

		fmt.Print("Tinggi Badan (cm): ")
		_, err = fmt.Scan(&tb)

		if err != nil || bb <= 0 || tb <= 0 {
			fmt.Println("Input berat tidak valid!")
			continue
		}
		break
	}

	bmi := bb / ((tb / 100) * (tb / 100))
	fmt.Printf("BMI Anda: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Kategori: Kurus")
	} else if bmi < 25 {
		fmt.Println("Kategori: Normal")
	} else if bmi < 30 {
		fmt.Println("Kategori: Gemuk")
	} else if bmi >= 30 {
		fmt.Println("Kategori: Obesitas")
	} else {
		fmt.Println("Input tidak valid!")
	}

}
