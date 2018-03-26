package main

import (
	"fmt"
)

func main() {
	a := "aaaa"
	b := "a*b*"
	fmt.Println(isMatch(a, b))
	a = "ab"
	b = ".*c"
	fmt.Println(isMatch(a, b))
	a = "aa"
	b = "a"
	fmt.Println(isMatch(a, b))
	a = "aa"
	b = ".*"
	fmt.Println(isMatch(a, b))
}

func printDP(dp [][]bool) {
	for i := range dp {
		fmt.Println(dp[i])
	}
}

func isMatch(s string, p string) bool {
	// dp[i][j]的含义是 p[0-i] 与 s[0-j]是否匹配。
	// 首先构建以 len(p)+1 行， len(s)+1 列的二维数组 多一行一列为最开始的空字符串准备
	// 初始化 0行 0列
	// 填表规则如下:
	// p[i-1] == s[j-1] : dp[i][j] = dp[i-1][j-1]
	// If p[i-1] == '.' : dp[i][j] = dp[i-1][j-1];
	// If p[i-1] == '*':
	// here are two sub conditions:
	// if p[i-2] != s[j-1] && p[i-2] != '.' : dp[i][j] = dp[i-2][j] //in this case, a* only counts as empty
	// if p[i-2] == s[j-1] or p[i-2] == '.':
	// 	dp[i][j] = dp[i][j-1] //in this case, a* counts as multiple a
	// 	dp[i][j] = dp[i-1][j] // in this case, a* counts as single a
	// 	dp[i][j] = dp[i-2][j] // in this case, a* counts as empty
	//
	//初始化dp辅助数组
	dp := [][]bool{}
	for i := 0; i < len(p)+1; i++ {
		dp = append(dp, make([]bool, len(s)+1))
	}
	dp[0][0] = true
	for i := 1; i < len(p)+1; i++ {
		dp[i][0] = i > 1 && p[i-1] == '*' && dp[i-2][0]
	}

	// 填表
	for i := 1; i < len(p)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if p[i-1] == '*' {
				if p[i-2] != s[j-1] && p[i-2] != '.' {
					dp[i][j] = dp[i-2][j]
				} else if p[i-2] == s[j-1] || p[i-2] == '.' {
					dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i-2][j]
				}
			} else {
				dp[i][j] = dp[i-1][j-1] && (p[i-1] == '.' || s[j-1] == p[i-1])
			}
		}
	}
	return dp[len(p)][len(s)]
}
