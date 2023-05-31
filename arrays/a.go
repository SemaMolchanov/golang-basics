package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)

	a := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}

	for j := size - 1; j >= 0; j-- {
		fmt.Print(a[j], " ")
	}

	fmt.Println()
}
