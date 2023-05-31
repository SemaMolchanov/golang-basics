package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println("Enter a string:")
	fmt.Scanln(&s)

	count := 1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && s[i-1] != ' ' {
			count++
		}
	}

	fmt.Println(count)
}
