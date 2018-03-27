package main

import "fmt"

func main() {
	a := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(a, 7))
	a = []int{1, 2}
	fmt.Println(combinationSum(a, 4))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	dfs(candidates, target, []int{}, &res, 0)
	return res
}

func dfs(candidates []int, target int, arr []int, res *[][]int, index int) {
	for i := index; i < len(candidates); i++ {
		newTarget := target - candidates[i]
		if newTarget > 0 {
			arr = append(arr, candidates[i])
			dfs(candidates, newTarget, arr, res, i)
			arr = (arr)[:len(arr)-1]
		} else if newTarget == 0 {
			arr = append(arr, candidates[i])
			*res = append(*res, append([]int{}, arr...))
			arr = (arr)[:len(arr)-1]
		}
	}
}
