package main

import "fmt"

func main() {
	a := "aa"
	b := "*"
	a = "zacabz"
	b = "*a?b*"
	a = "aa"
	b = "aa"
	fmt.Println(isMatch(a, b))
}

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(s) != 0 && len(p) == 0 {
		return false
	} else if len(s) == 0 && p[0] != '*' {
		return false
	}
	sI := 0
	pI := 0
	for sI < len(s) {
		fmt.Println("------")
		fmt.Println(s[sI:])
		fmt.Println(p[pI:])
		if pI == len(p) {
			return false
		}
		if p[pI] == '?' || s[sI] == p[pI] {
			sI++
			pI++
			continue
		} else if p[pI] == '*' {
			fmt.Println(">>>")
			if isMatch(s[sI+1:], p[pI+1:]) {
				return true
			}
			sI++
			fmt.Println("<<<")
		} else if s[sI] != p[pI] {
			return false
		}
	}
	if len(p) != 0 && isMatch(s[sI:], p[pI+1:]) {
		return true
	}
	// if len(p)-pI > 1 || p[pI] != '*' {
	// 	return false
	// }
	return false
}
