package main

import (
	"fmt"
)

func main() {
	test([]int{}, 0, []int{})
	test([]int{1}, 1, []int{1})
	test([]int{1, 1, 2}, 3, []int{1, 1, 2})
	test([]int{1, 1, 1, 2, 2}, 4, []int{1, 1, 2, 2})
	test([]int{1, 1, 1, 2, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3})
	test([]int{1, 1, 1, 1, 3, 3}, 4, []int{1, 1, 3, 3})
}

func test(nums []int, count int, r []int) {
	if count != removeDuplicates(nums) {
		fmt.Println(nums, count, r)
	}
	for i := 0; i < count; i++ {
		if nums[i] != r[i] {
			fmt.Println(nums, count, r)
		}
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	res := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			if count <= 2 {
				nums[res] = nums[i]
				res++
			}
		} else {
			count = 1
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
