package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scanf("%s %d %f\n", &prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println()
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := int((tumbuh[i] + 1) * float64(pop[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	InputData(&prov, &pop, &tumbuh)
	idxTercepat := ProvinsiTercepat(tumbuh)
	fmt.Println(prov[idxTercepat])
	fmt.Println()
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[idxTercepat])
	fmt.Println()
	var namaCari string
	fmt.Scanf("%s", &namaCari)
	idxCari := IndeksProvinsi(prov, namaCari)
	fmt.Println("Data provinsi yang dicari :", prov[idxCari])
	Prediksi(prov, pop, tumbuh)
}
