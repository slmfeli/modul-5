package main

import (
	"fmt"
)

func printStars(n int) {
	if n <= 0 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printStars(n - 1)
}

func main() {
	var n int
	fmt.Print("Masukkan untuk menampilkan pola bintang: ")
	fmt.Scan(&n)

	printStars(n)
}