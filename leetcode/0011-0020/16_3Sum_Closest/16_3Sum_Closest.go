package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{0, 0, 0}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	min := math.MaxInt32
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(min)) {
				min = target - sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return target
			}
		}
	}
	return target - min
}
