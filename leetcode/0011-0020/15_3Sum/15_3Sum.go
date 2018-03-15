package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))
	a = []int{1, 1, -2}
	fmt.Println(threeSum(a))
	a = []int{0, 0, 0}
	fmt.Println(threeSum(a))
}

func threeSum(nums []int) [][]int {
	rArr := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				rArr = append(rArr, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
		}
	}
	return rArr
}
