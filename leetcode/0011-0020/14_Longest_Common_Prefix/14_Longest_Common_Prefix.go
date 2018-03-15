package main

import (
	"fmt"
)

func main() {
	a := []string{}
	a = []string{
		"",
		"12A",
	}
	fmt.Println(longestCommonPrefix(a))
	fmt.Println("adsfas")
	a = []string{
		"c",
		"c",
	}
	fmt.Println(longestCommonPrefix(a))
	fmt.Println("adsfas")
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	index := 0
	for true {
		var i int
		var b byte
		if index < len(strs[0]) {
			b = strs[0][index]
		}
		for i = 1; i < len(strs); i++ {
			if index < len(strs[i]) {
				if b == strs[i][index] {
					continue
				} else {
					break
				}
			} else {
				break
			}
		}
		if i == len(strs) {
			index++
		} else {
			break
		}
	}
	return strs[0][0:index]
}
