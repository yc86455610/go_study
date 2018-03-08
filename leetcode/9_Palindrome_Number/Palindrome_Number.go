package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	var reverseInteger int
	for x > reverseInteger {
		reverseInteger = reverseInteger*10 + x%10
		x = x / 10
	}

	return x == reverseInteger || x == reverseInteger/10
}

func main() {
	a := 121
	fmt.Println(isPalindrome(a))
}
