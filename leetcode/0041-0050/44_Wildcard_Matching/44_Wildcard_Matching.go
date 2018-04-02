package main

import "fmt"

func main() {
	a := "aa"
	b := "*"
	fmt.Println(isMatch(a, b))
	a = "zacabz"
	b = "*a?b*"
	fmt.Println(isMatch(a, b))
	a = "aa"
	b = "a"
	fmt.Println(isMatch(a, b))
	a = "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaabaa"
	b = "a*******b"
	fmt.Println(isMatch(a, b))
	a = "babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb"
	b = "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a"
	fmt.Println(isMatch(a, b))
}

func isMatch(s string, p string) bool {
	star := -1
	sStation := -1
	i, j := 0, 0
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && (p[j] == '*') {
			star = j
			j++
			sStation = i
		} else if star != -1 {
			j = star + 1
			sStation++
			i = sStation
		} else {
			return false
		}
	}
	for j < len(p) {
		if p[j] == '*' {
			j++
		} else {
			return false
		}
	}
	return true
}
