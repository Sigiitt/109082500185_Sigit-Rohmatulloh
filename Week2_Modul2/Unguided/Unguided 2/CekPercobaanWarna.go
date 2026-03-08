package main

import "fmt"

func main() {
	berhasil := true

	for i := 1; i <= 5; i++ {
		var g1, g2, g3, g4 string
		fmt.Printf("Percobaan %d : ", i)
		fmt.Scan(&g1, &g2, &g3, &g4)

		if g1 != "merah" || g2 != "kuning" || g3 != "hijau" || g4 != "ungu" {
			berhasil = false
		}
	}

	if berhasil {
		fmt.Println("BERHASIL :  true")
	} else {
		fmt.Println("BERHASIL :  false")
	}
}
