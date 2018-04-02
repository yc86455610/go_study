package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 1, 2, 2, 3}
	fmt.Println(permuteUnique(a))
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums []int, arr []int, res *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		} else {
			temp := make([]int, len(nums))
			copy(temp, nums)
			dfs(append(temp[:i], temp[i+1:]...), append(arr, nums[i]), res)
		}
	}
}
