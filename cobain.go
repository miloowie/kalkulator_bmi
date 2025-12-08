package main

import "fmt"

func main() {
	var bb, tb float64
	var gender string
	var bersedia bool
	var test int
	fmt.Println(test)

	fmt.Println("Kalkulator BMI")
	for {
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
			_, err1 := fmt.Scan(&bb)

			fmt.Print("Tinggi Badan (cm): ")
			_, err2 := fmt.Scan(&tb)

			if err1 != nil || err2 != nil || bb <= 0 || tb <= 0 {
				fmt.Println("Input berat/tinggi tidak valid!")
				continue
			}
			break
		}

		for {
			fmt.Print("Apakah anda bersedia melanjutkan perhitungan BMI? (true / false) :")
			fmt.Scan(&bersedia)

			if bersedia == true || bersedia == false {
				break
			} else if bersedia != true && bersedia != false {
				fmt.Println("Input tidak valid! Mohon hanya masukkan true atau false")
			}
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

		var ulang string
		for {
			fmt.Print("Hitung lagi? (y/n): ")
			fmt.Scan(&ulang)
			if ulang == "y" {
				main()
				return
			} else if ulang == "n" {
				fmt.Println("Program selesai.")
				return
			}

			fmt.Println("Input tidak valid! Masukkan hanya y atau n.")
		}

	}
	fmt.Println()
}
