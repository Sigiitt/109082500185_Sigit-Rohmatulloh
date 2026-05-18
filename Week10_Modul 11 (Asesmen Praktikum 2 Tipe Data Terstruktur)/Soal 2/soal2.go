package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(arr *arrayMahasiswa, n *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&arr[i].NIM, &arr[i].nama, &arr[i].nilai)
	}
}

func cariNilaiPertama(arr arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			return arr[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(arr arrayMahasiswa, n int, nim string) int {
	max := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			if arr[i].nilai > max {
				max = arr[i].nilai
			}
		}
	}
	return max
}

func main() {
	var arr arrayMahasiswa
	var n int
	inputData(&arr, &n)
	var nimCari string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)
	pertama := cariNilaiPertama(arr, n, nimCari)
	terbesar := cariNilaiTerbesar(arr, n, nimCari)
	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, pertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, terbesar)
}
