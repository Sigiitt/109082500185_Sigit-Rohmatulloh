package main

import "fmt"

// Alias tipe data suhu bertipe float64 (dideklarasikan global)
type suhu float64

// Fungsi konversi Celsius ke Reamur
func CelciusToReamur(Celcius suhu) suhu {
	return Celcius * 4 / 5
}

// Fungsi konversi Celsius ke Fahrenheit
func CelciusToFahrenheit(Celcius suhu) suhu {
	return Celcius*9/5 + 32
}

// Fungsi konversi Celsius ke Kelvin
func CelciusToKelvin(Celcius suhu) suhu {
	return Celcius + 273.15
}

func main() {
	var input suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&input)

	fmt.Println()
	fmt.Printf("%.4g celcius = %.4g reamur\n", input, CelciusToReamur(input))
	fmt.Printf("%.4g celcius = %.4g fahrenheit\n", input, CelciusToFahrenheit(input))
	fmt.Printf("%.4g celcius = %.4g kelvin\n", input, CelciusToKelvin(input))
}
