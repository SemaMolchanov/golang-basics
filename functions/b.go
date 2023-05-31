package main

import "fmt"

func negativePower(a float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == -1 {
		return 1 / a
	}
	return (1 / a) * negativePower(a, n+1)
}

func power(a float64, n int) float64 {
	if n < 0 {
		return negativePower(a, n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	return a * power(a, n-1)
}

func main() {
	var a float64
	var n int
	fmt.Scan(&a, &n)
	fmt.Println(power(a, n))
}
