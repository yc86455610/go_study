package main

import (
	"fmt"
	"math"
)

func main() {
	var s string
	s = "babad"
	fmt.Println(longestPalindrome(s))
}

func preprocessor(s string) []byte {
	var newByte []byte
	newByte = append(newByte, '$')
	for i := 0; i < len(s); i++ {
		newByte = append(newByte, '#')
		newByte = append(newByte, s[i])
	}
	newByte = append(newByte, '#')
	newByte = append(newByte, '@')

	return newByte
}

func postprocessing(b []byte) string {
	var nb []byte
	for i := 0; i < len(b); i++ {
		if b[i] == '#' {
			continue
		}
		nb = append(nb, b[i])
	}
	return string(nb)
}

func longestPalindrome(s string) string {
	var b []byte
	b = preprocessor(s)

	var p []int = make([]int, len(b))
	var maxid int = 0
	var id int = 0
	var result int = 0
	var position int = 0
	for i := 1; i < len(b)-1; i++ {
		if maxid > i {
			p[i] = int(math.Min(float64(p[2*id-i]), float64(maxid-i)))
		} else {
			p[i] = 1
		}
		for b[i-p[i]] == b[i+p[i]] {
			p[i]++
		}
		if maxid < i+p[i] {
			maxid = i + p[i]
			id = i
		}
		if result < p[i]-1 {
			result = p[i] - 1
			position = i
		}
	}
	r := postprocessing(b[position-result : position+result+1])

	return string(r)
}
