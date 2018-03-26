package main

import "fmt"

func main() {
	a := "hello"
	b := "ll"
	fmt.Println(strStr(a, b))
	a = "aaa"
	b = "aa"
	fmt.Println(strStr(a, b))
	a = "mississippi"
	b = "issipi"
	fmt.Println(strStr(a, b))
}
func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func myStrStr(haystack string, needle string) int {
	index := -1
	if len(haystack) < len(needle) {
		return index
	}
	for i := 0; i < len(haystack)-len(needle); i++ {
		var hi, j int
		for hi, j = i, 0; j < len(needle); j++ {
			if haystack[hi] == needle[j] {
				hi++
			} else {
				break
			}
		}
		if j == len(needle) {
			index = i
			break
		}
	}
	return index
}
