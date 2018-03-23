package main

import (
	"fmt"
)

func main() {
	a := ""
	fmt.Println(longestValidParentheses(a))
	a = "(()())"
	fmt.Println(longestValidParentheses(a))
}

func longestValidParentheses(s string) int {
	s = ")" + s
	dp := []int{0}
	maxLen := 0
	for i := 1; i < len(s); i++ {
		dp = append(dp, 0)
		if s[i] == ')' {
			if s[i-1-dp[i-1]] == '(' {
				dp[i] = dp[i-1] + 2
			}
			dp[i] += dp[i-dp[i]]
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
