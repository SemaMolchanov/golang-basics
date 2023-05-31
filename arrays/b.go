package main

import (
	"fmt"
)

func main() {
	var size, a, b, c, d int
	fmt.Scan(&size, &a, &b, &c, &d)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}

	reverse(arr[a-1 : b])
	reverse(arr[c-1 : d])

	for i := 0; i < size; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
