package main

import "fmt"

func star(s string, i int) {
	if i == len(s)-1 {
		fmt.Print(string(s[i]))
		return
	}
	fmt.Print(string(s[i]) + "*")
	star(s, i+1)
}

func main() {
	var s string
	fmt.Scan(&s)
	star(s, 0)
}
