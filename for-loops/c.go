package main

import "fmt"

func main() {
	var n int
	var sum int = 0
	fmt.Scanln(&n)
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	fmt.Println(sum)
}
