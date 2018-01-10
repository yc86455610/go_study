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
	var nb []byte
	nb = append(nb, '$')
	for i := 0; i < len(s); i++ {
		nb = append(nb, '#')
		nb = append(nb, s[i])
	}
	nb = append(nb, '#')
	nb = append(nb, '@')

	return nb
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
	var b []byte = preprocessor(s)
	var p []int = make([]int, len(b))
	var maxid, id, result, position int
	var r string

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
	r = postprocessing(b[position-result : position+result+1])

	return r
}
