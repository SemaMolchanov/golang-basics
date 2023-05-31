package main

import (
	"fmt"
	"unicode"
)

func changeRegister(c *rune) rune {
	if unicode.IsLower(*c) {
		return unicode.ToUpper(*c)
	} else if unicode.IsUpper(*c) {
		return unicode.ToLower(*c)
	}
	return *c
}

func main() {
	var c rune
	fmt.Scanf("%c", &c)
	fmt.Printf("%c", changeRegister(&c))
}
