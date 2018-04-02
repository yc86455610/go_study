package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(permute(a))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, []int{}, &res)
	return res
}
func dfs(nums []int, arr []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, arr)
	}
	for i := 0; i < len(nums); i++ {
		temp := copyIntSlice(nums)
		dfs(append(temp[:i], temp[i+1:]...), append(arr, nums[i]), res)
	}
}

func copyIntSlice(a []int) []int {
	res := make([]int, len(a))
	copy(res, a)
	fmt.Println(res)
	return res
}
