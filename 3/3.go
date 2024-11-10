
package main

import (
	"fmt"
)

func printFactors(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ") 
	}
	printFactors(n, i+1) 
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n) 

	fmt.Printf("Faktor dari %d: ", n)
	printFactors(n, 1) 
	fmt.Println() 
}
