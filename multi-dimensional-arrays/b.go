package main

import (
	"fmt"
)

func main() {
	var row, col int
	fmt.Scan(&row, &col)

	input := make([][]int, row)
	turned := make([][]int, col)
	for i := 0; i < row; i++ {
		input[i] = make([]int, col)
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			fmt.Scan(&input[r][c])
		}
	}

	for r := 0; r < row; r++ {
		turned[r] = make([]int, row)
		for c := 0; c < col; c++ {
			turned[c][r] = input[row-r-1][c]
		}
	}

	fmt.Println(col, row)
	for r := 0; r < col; r++ {
		for c := 0; c < row; c++ {
			fmt.Print(turned[r][c], " ")
		}
		fmt.Println()
	}
}
