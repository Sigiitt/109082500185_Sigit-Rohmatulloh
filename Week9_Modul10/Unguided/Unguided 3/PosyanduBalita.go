package main

import "fmt"

const NMAX = 100

type arrBalita [NMAX]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	j := 1
	for j < n {
		if arrBerat[j] < *bMin {
			*bMin = arrBerat[j]
		}
		if arrBerat[j] > *bMax {
			*bMax = arrBerat[j]
		}
		j = j + 1
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	sum := 0.0
	i := 0
	for i < n {
		sum = sum + arrBerat[i]
		i = i + 1
	}
	return sum / float64(n)
}

func main() {
	var n int
	var arr arrBalita
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)
	i := 0
	for i < n {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
		i = i + 1
	}
	var bMin, bMax float64
	hitungMinMax(arr, n, &bMin, &bMax)
	avg := rerata(arr, n)
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", avg)
}
