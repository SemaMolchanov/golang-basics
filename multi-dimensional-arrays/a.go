package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)

	m := make([][]int, size)
	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
	}

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			fmt.Scan(&m[r][c])
		}
	}

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if m[r][c] != m[c][r] {
				fmt.Println("no")
				return
			}
		}
	}

	fmt.Println("yes")
}
