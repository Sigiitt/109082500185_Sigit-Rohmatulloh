package main

import "fmt"

func main() {
	var n int
	// Input bilangan bulat positif N
	fmt.Scan(&n)

	// Memanggil fungsi rekursif dimulai dari angka 1
	cetakGanjil(n, 1)
	fmt.Println()
}

func cetakGanjil(n int, i int) {
	// Base-case: Berhenti jika angka i sudah melewati N
	if i <= n {
		// Cek apakah angka ganjil
		if i%2 != 0 {
			fmt.Print(i, " ")
		}

		// Recursive-case: Panggil fungsi untuk angka berikutnya
		cetakGanjil(n, i+1)
	}
}
