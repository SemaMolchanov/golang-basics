package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var i int = 1
	for i <= n {
		if i*i > n {
			break
		}
		println(i * i)
		i++
	}
}
