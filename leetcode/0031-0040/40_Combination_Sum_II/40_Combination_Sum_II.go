package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(a, 8))
	a = []int{1, 2}
	fmt.Println(combinationSum2(a, 4))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	dfs(candidates, target, []int{}, &res, 0)
	return res
}

func dfs(candidates []int, target int, arr []int, res *[][]int, index int) {
	for i := index; i < len(candidates); i++ {
		newTarget := target - candidates[i]
		arr = append(arr, candidates[i])
		if newTarget > 0 {
			dfs(candidates, newTarget, arr, res, i+1)
		} else if newTarget == 0 && !inRange(arr, *res) {
			*res = append(*res, append([]int{}, arr...))
		}
		arr = (arr)[:len(arr)-1]
	}
}
func inRange(a []int, r [][]int) bool {
	for i := 0; i < len(r); i++ {
		if equal(a, r[i]) {
			return true
		}
	}
	return false
}
func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}
