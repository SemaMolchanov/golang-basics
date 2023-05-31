package main

import "fmt"

func main() {
	var n int
	var factorial int = 1
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if factorial == 0 {
			factorial = 1
			break
		}
		factorial *= i
	}
	fmt.Println(factorial)
}
