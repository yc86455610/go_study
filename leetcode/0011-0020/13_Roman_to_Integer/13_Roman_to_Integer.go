package main

import (
	"fmt"
)

func main() {
	a := "DCXXI"
	fmt.Println(romanToInt(a))
	a = "I"
	fmt.Println(romanToInt(a))
}

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	rInt := m[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if m[s[i]] < m[s[i+1]] {
			rInt -= m[s[i]]
		} else {
			rInt += m[s[i]]
		}
	}
	return rInt
}
