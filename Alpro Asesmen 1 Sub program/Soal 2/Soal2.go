package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	var max int
	if tujuan == "domestik" {
		max = 3
	} else {
		max = 8
	}

	if jumlahHari > max {
		return max
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	return 620000 * jumlahMhs
}

func perhitunganBiaya(jumlahMhs int, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)
	biayaDasar := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaDasar)
	} else {
		*totalBiaya = float64(hariDitanggung) * (float64(biayaDasar) * 1.5)
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
