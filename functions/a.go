package main

import "fmt"

func exclusiveOr(a, b bool) bool {
	return a != b
}

func main() {
	var a, b bool
	fmt.Scan(&a, &b)
	fmt.Println(exclusiveOr(a, b))
}
