package main

import "fmt"

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func main() {
	var year int
	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println("Kabisat : true")
	} else {
		fmt.Println("Kabisat : false")
	}
}
