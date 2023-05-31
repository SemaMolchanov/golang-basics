package main

import "fmt"

func main() {
	var n int
	var sum, count float64 = 0, 0
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		sum += float64(n)
		count++
	}
	fmt.Println(sum / count)
}
