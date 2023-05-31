package main

import "fmt"

func main() {
	var n int
	var sum int = 0
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	fmt.Println(sum)
}
