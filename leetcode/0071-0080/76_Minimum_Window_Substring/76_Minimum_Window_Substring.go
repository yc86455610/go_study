package main

import (
	"fmt"
	"math"
)

func main() {
	test(1, "ADOBECODEBANC", "ABC", "BANC")
	test(2, "A", "AA", "")
}

func test(n int, s, t, r string) {
	if r != minWindow(s, t) {
		fmt.Println(n, s, t, r)
	}
}

func minWindow(s string, t string) string {
	m := map[byte]int{}
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	count := 0
	left := 0
	right := 0
	minStart := 0
	minEnd := math.MaxInt64
	for right <= len(s) {
		if count >= len(t) {
			if right-left < minEnd-minStart {
				minEnd = right
				minStart = left
			}
			if temp, ok := m[s[left]]; ok {
				if temp == 0 {
					count--
				}
				m[s[left]]++
			}
			left++
		} else {
			if right < len(s) {
				if temp, ok := m[s[right]]; ok {
					if temp > 0 {
						count++
					}
					m[s[right]]--
				}
			}
			right++
		}

	}
	if minEnd-minStart > len(s) {
		return ""
	}
	return s[minStart:minEnd]
}
