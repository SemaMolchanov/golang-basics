package main

import (
	"fmt"
)

func areEqual(s1, s2 string) bool {
	return s1 == s2
}

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	if areEqual(s1, s2) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
