package main

import "fmt"

func main() {
	var suara, totalMasuk, suaraSah int
	var hitungSuara [21]int

	totalMasuk = 0
	suaraSah = 0

	// Membaca input sampai angka 0
	fmt.Scan(&suara)
	for suara != 0 {
		totalMasuk++
		if suara >= 1 && suara <= 20 {
			hitungSuara[suara]++
			suaraSah++
		}
		fmt.Scan(&suara)
	}

	// Logika mencari Ketua (Terbanyak 1) dan Wakil (Terbanyak 2)
	ketua := -1
	wakil := -1

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			// Jika calon i punya suara lebih banyak dari ketua saat ini
			if ketua == -1 || hitungSuara[i] > hitungSuara[ketua] {
				wakil = ketua // Ketua lama turun jadi wakil
				ketua = i     // Calon i jadi ketua baru
			} else if wakil == -1 || hitungSuara[i] > hitungSuara[wakil] {
				// Jika suara i lebih banyak dari wakil tapi tidak lebih dari ketua
				wakil = i
			}
			// Catatan: Karena loop berjalan dari i=1 ke 20,
			// kondisi nomor terkecil sudah otomatis terpenuhi (siapa yang duluan dicek).
		}
	}

	// Menampilkan hasil
	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	if ketua != -1 {
		fmt.Printf("Ketua RT: %d\n", ketua)
	}
	if wakil != -1 {
		fmt.Printf("Wakil ketua: %d\n", wakil)
	}
}
