package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)

	a := make([]int, 35)

	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}

	tmpr := a[0]

	for i := 1; i < size; i++ {
		tmpr, a[i] = a[i], tmpr
	}

	tmpr, a[0] = a[0], tmpr

	for i := 0; i < size; i++ {
		fmt.Print(a[i], " ")
	}

	fmt.Println()
}
