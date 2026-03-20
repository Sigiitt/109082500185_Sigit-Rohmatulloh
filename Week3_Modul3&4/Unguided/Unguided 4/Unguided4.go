package main

import "fmt"

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Println("Luas Persegi :", luas)
	fmt.Println("Keliling Persegi :", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println("Luas Persegi Panjang :", luas)
	fmt.Println("Keliling Persegi Panjang :", keliling)
}

func hitungLingkaran(jarijari float64) {
	const phi = 3.14

	luas := phi * jarijari * jarijari
	keliling := 2 * phi * jarijari

	fmt.Println("Luas Lingkaran :", luas)
	fmt.Println("Keliling Lingkaran :", keliling)
}

func main() {

	var pilihan int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:
		var sisi int
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)

	case 2:
		var panjang, lebar int
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)

	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&r)
		hitungLingkaran(r)

	default:
		fmt.Println("Pilihan tidak tersedia")
	}

}
