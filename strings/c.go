package main

import (
	"fmt"
)

func main() {
	var s, maxs string
	for {
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		if len(s) > len(maxs) {
			maxs = s
		}
	}
	fmt.Println(maxs)
	fmt.Println(len(maxs))
}
