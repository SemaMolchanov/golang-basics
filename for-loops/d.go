package main

import "fmt"

func main() {
	var n int
	var count int = 0
	fmt.Scanln(&n)
	for i := n; i > 0; i /= 10 {
		count++
	}
	var max int = count
	for ; n > 0; n /= 10 {
		var digit int = n % 10
		if count == max && digit == 0 {
			continue
		}
		fmt.Print(digit)
		count--
	}
	fmt.Println(count)
}
