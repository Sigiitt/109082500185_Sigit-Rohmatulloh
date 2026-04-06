package main

import "fmt"

func main() {
	var n int
	// Input suku ke-n yang ingin dicari
	fmt.Scan(&n)
	// Menampilkan hasil suku ke-n dari deret fibonacci
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	// Base-case: suku ke-0 adalah 0, suku ke-1 adalah 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		// Recursive-case: Sn = Sn-1 + Sn-2
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
