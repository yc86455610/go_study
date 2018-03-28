package main

import (
	"fmt"
)

func main() {
	a := []int{1, 44, 2, 0, 55}
	fmt.Println(firstMissingPositive(a))
}

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] <= 0 || nums[i] == i+1 || nums[i] >= len(nums) || nums[i] == nums[nums[i]-1] {
			i++
			continue
		}
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
