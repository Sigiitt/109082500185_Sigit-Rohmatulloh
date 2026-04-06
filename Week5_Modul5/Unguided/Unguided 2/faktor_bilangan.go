package main

import "fmt"

func main() {
	var n int
	// Input bilangan bulat positif N
	fmt.Scan(&n)
	// Memanggil fungsi rekursif dimulai dari pembagi 1
	cariFaktor(n, 1)
	fmt.Println() // Baris baru di akhir
}

func cariFaktor(n int, i int) {
	// Base-case: jika pembagi i sudah melebihi n, berhenti
	if i <= n {
		// Jika n habis dibagi i, maka i adalah faktor
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		// Recursive-case: panggil kembali dengan pembagi selanjutnya (i + 1)
		cariFaktor(n, i+1)
	}
}
