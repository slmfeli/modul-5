package main

import (
	"fmt"
)

func printOddNumbers(n int, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	printOddNumbers(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	fmt.Printf("Barisan bilangan ganjil dari 1 hingga %d: ", n)
	printOddNumbers(n, 1)
	fmt.Println()
}