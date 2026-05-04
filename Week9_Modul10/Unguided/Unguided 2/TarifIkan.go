package main

import "fmt"

const NMAX = 1000

type arrIkan [NMAX]float64

func main() {
	var x, y int
	var arr arrIkan
	var totalWadah [NMAX]float64
	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah : ")
	fmt.Scan(&x, &y)
	i := 0
	for i < x {
		fmt.Printf("Masukkan berat ikan ke-%d : ", i+1)
		fmt.Scan(&arr[i])
		i = i + 1
	}
	numWadah := x / y
	if x%y != 0 {
		numWadah = numWadah + 1
	}
	i = 0
	for i < x {
		wadahIdx := i / y
		totalWadah[wadahIdx] = totalWadah[wadahIdx] + arr[i]
		i = i + 1
	}
	w := 0
	for w < numWadah {
		if w > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", totalWadah[w])
		w = w + 1
	}
	fmt.Println()
	sum := 0.0
	i = 0
	for i < x {
		sum = sum + arr[i]
		i = i + 1
	}
	fmt.Printf("%.2f\n", sum/float64(x))
}
