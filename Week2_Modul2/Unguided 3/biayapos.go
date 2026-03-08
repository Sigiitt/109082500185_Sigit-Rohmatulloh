package main

import "fmt"

func main() {
	var totalGram int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&totalGram)

	kg := totalGram / 1000
	sisa := totalGram % 1000

	biayaKg := kg * 10000

	var biayaSisa int
	if sisa > 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	// Jika total berat lebih dari 10kg, sisa berat digratiskan
	if kg > 10 {
		biayaSisa = 0
	}

	totalBiaya := biayaKg + biayaSisa

	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat  : %d kg + %d gram\n", kg, sisa)
	fmt.Printf("Detail biaya  : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
