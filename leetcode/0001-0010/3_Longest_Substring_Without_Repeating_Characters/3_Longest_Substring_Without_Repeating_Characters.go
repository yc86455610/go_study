package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	head := 0
	tail := 0
	maxcount := 0
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; v >= head && ok {
			head = v + 1
		}
		m[s[i]] = i
		tail = i
		if tail-head+1 > maxcount {
			maxcount = tail - head + 1
		}
		// fmt.Println(i, string(s[i]), head, tail, maxcount, m)
	}
	return maxcount
}

func main() {
	a := "tmmzuxt"
	b := "bbbb"
	c := "pwwkew"
	count := lengthOfLongestSubstring(a)
	fmt.Println(count)
	count = lengthOfLongestSubstring(b)
	fmt.Println(count)
	count = lengthOfLongestSubstring(c)
	fmt.Println(count)
}
