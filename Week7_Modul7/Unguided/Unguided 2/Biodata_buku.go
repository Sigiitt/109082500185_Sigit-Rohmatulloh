package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Alias tipe data
type angka int
type kata string

// Struct Buku
type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buku Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Masukkan judul buku : ")
	judul, _ := reader.ReadString('\n')
	buku.judul = kata(strings.TrimSpace(judul))

	fmt.Print("Masukkan nama penulis : ")
	penulis, _ := reader.ReadString('\n')
	buku.penulis = kata(strings.TrimSpace(penulis))

	fmt.Print("Masukkan penerbit : ")
	penerbit, _ := reader.ReadString('\n')
	buku.penerbit = kata(strings.TrimSpace(penerbit))

	var tahun, halaman int
	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scan(&tahun)
	buku.tahunTerbit = angka(tahun)

	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&halaman)
	buku.jumlahHalaman = angka(halaman)

	fmt.Println()
	fmt.Println("=== BIODATA BUKU ===")
	fmt.Printf("Judul Buku    : %v\n", buku.judul)
	fmt.Printf("Penulis       : %v\n", buku.penulis)
	fmt.Printf("Penerbit      : %v\n", buku.penerbit)
	fmt.Printf("Tahun Terbit  : %v\n", buku.tahunTerbit)
	fmt.Printf("Jumlah Halaman: %v\n", buku.jumlahHalaman)
}
