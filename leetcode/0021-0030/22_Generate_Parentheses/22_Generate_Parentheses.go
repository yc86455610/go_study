package main

import (
	"fmt"
)

func main() {
	a := 3
	fmt.Println(generateParenthesis(a))
}

func dfs(rArr []string, s string, left int, right int) []string {
	if left == 0 && right == 0 {
		rArr = append(rArr, s)
		return rArr
	}
	if left > 0 {
		rArr = dfs(rArr, s+"(", left-1, right)
	}
	if right > 0 && left < right {
		rArr = dfs(rArr, s+")", left, right-1)
	}
	return rArr
}

func generateParenthesis(n int) []string {
	return dfs([]string{}, "", n, n)
}
