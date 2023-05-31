package main

import (
	"fmt"
)

func isPalindrome(res string) bool {
	for i := 0; i <= len(res)/2; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var res, s string
	for {
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		res += s
	}
	if isPalindrome(res) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
