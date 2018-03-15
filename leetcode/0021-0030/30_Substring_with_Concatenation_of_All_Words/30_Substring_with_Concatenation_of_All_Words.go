package main

import (
	"fmt"
)

func main() {
	a := "barfoothefoobarman"
	b := []string{"foo", "bar"}
	fmt.Println(findSubstring(a, b))
	a = "wordgoodgoodgoodbestword"
	b = []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(a, b))
	a = "aaabbbc"
	b = []string{"a", "a", "b", "b", "c"}
	fmt.Println(findSubstring(a, b))
	a = "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	b = []string{"fooo", "barr", "wing", "ding", "wing"}
	fmt.Println(findSubstring(a, b))
	a = "barfoofoobarthefoobarman"
	b = []string{"bar", "foo", "the"}
	fmt.Println(findSubstring(a, b))
}

func findSubstring(s string, words []string) []int {
	rArr := []int{}

	lWord := len(words[0])
	if len(s) == 0 || len(s) < len(words)*len(words[0]) {
		return rArr
	}
	m := map[string]int{}
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	for i := 0; i <= len(s)-len(words)*lWord; i++ {
		var j int
		tmp := map[string]int{}
		for j = 0; j < len(words); j++ {
			subStr := s[i+j*lWord : i+(j+1)*lWord]
			tmp[subStr]++
			if tmp[subStr] > m[subStr] {
				break
			}
		}
		if j == len(words) {
			rArr = append(rArr, i)
		}
	}
	return rArr
}
