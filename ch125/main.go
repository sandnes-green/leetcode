package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// buf := []byte(s)
	str := strings.ToLower(s)
	l, h := 0, len(s)-1

	for l < h {
		for !((str[l] > 47 && str[l] < 58) || (str[l] > 96 && str[l] < 123)) && l < h {
			l++
		}
		for !((str[h] > 47 && str[h] < 58) || (str[h] > 96 && str[h] < 123)) && l < h {
			h--
		}
		if l < h && str[l] != str[h] {
			return false
		}
		l++
		h--
	}
	return true
}

func main() {
	s := ".,"

	fmt.Println(isPalindrome(s))
}
