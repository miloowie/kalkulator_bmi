package main

import "fmt"

func main() {
	var bb, tb float64
	var gender string

	fmt.Println("Kalkulator BMI")

	for {
		for {
			fmt.Print("Jenis Kelamin (L/P): ")
			fmt.Scan(&gender)

			if gender == "L" || gender == "P" {
				break
			} else {
				fmt.Println("Input tidak valid! Mohon masukkan L atau P")
			}
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
		} else {
			fmt.Println("Kategori: Obesitas")
		}

		for {
			var ulang string
			fmt.Print("Hitung lagi? (y/n): ")
			fmt.Scan(&ulang)

			if ulang == "y" {
				main()
				return
			} else if ulang == "n" {
				fmt.Println("Program selesai, jangan lupa menjalani hidup sehat!.")
				return
			}

			fmt.Println("Input tidak valid! Mohon masukkan hanya y atau n.")
		}
		fmt.Println()
	}
}
