package main

import "fmt"

const NMAX = 1000

type arrKelinci [NMAX]float64

func findMinMax(arr arrKelinci, n int) (float64, float64) {
	min := arr[0]
	max := arr[0]
	j := 1
	for j < n {
		if arr[j] < min {
			min = arr[j]
		}
		if arr[j] > max {
			max = arr[j]
		}
		j = j + 1
	}
	return min, max
}

func main() {
	var n int
	var arr arrKelinci
	fmt.Print("Masukkan jumlah kelinci : ")
	fmt.Scan(&n)
	i := 0
	for i < n {
		fmt.Printf("Masukkan berat kelinci ke-%d : ", i+1)
		fmt.Scan(&arr[i])
		i = i + 1
	}
	min, max := findMinMax(arr, n)
	fmt.Printf("Berat terkecil : %.2f\n", min)
	fmt.Printf("Berat terbesar : %.2f\n", max)
}
