package main

import "fmt"

func main() {
	a := "func length  "
	fmt.Println(lengthOfLastWord(a))
}

func lengthOfLastWord(s string) int {
	sign := false
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if sign {
				break
			} else {
				continue
			}
		} else {
			sign = true
			count++
		}
	}

	return count
}
