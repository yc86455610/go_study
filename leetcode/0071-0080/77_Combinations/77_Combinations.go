package main

import (
	"fmt"
)

func main() {
	test(4, 2, [][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4}})
}

func test(n int, k int, r [][]int) {
	t := combine(n, k)
	fmt.Println(t)
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[0]); j++ {
			if r[i][j] != t[i][j] {
				fmt.Println(n, k, r)
			}
		}
	}
}

func combine(n int, k int) [][]int {
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, i+1)
	}
	res := [][]int{}
	dfs(a, k, &[]int{}, &res)
	return res
}
func dfs(a []int, k int, temp *[]int, res *[][]int) {
	if k == 0 {
		c := make([]int, len(*temp))
		copy(c, *temp)
		*res = append(*res, c)
		return
	}
	for i := 0; i < len(a); i++ {
		t := make([]int, len(a))
		copy(t, a)
		*temp = append(*temp, a[i])
		t = append([]int{}, t[i+1:]...)
		dfs(t, k-1, temp, res)
		*temp = (*temp)[:len(*temp)-1]
	}
}
