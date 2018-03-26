package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(a, 0))
	a = []int{1, -2, -5, -4, -3, 3, 3, 5}
	fmt.Println(fourSum(a, -11))
}

func fourSum(nums []int, target int) [][]int {
	rArr := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					rArr = append(rArr, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else if sum > target {
					right--
				}
			}
		}
	}
	return rArr
}
