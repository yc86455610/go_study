package main

import "fmt"

func main() {
	test(1, 1)
	test(2, 2)
	test(3, 3)
}

func test(x, y int) {
	if y != climbStairs(x) {
		fmt.Printf("wrong answer. x: %d, climbStairs:%d, expacted: %d\n", x, climbStairs(x), y)
	}
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	count := make([]int, n+1)
	count[0] = 1
	count[1] = 1
	count[2] = 2
	for i := 3; i <= n; i++ {
		count[i] = count[i-1] + count[i-2]
	}
	return count[n]
}
