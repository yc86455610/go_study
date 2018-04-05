package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a))
	a = []int{-2, -1}
	fmt.Println(maxSubArray(a))
	a = []int{1, 2}
	fmt.Println(maxSubArray(a))
}

func maxSubArray(nums []int) int {
	res := math.MinInt32
	max := 0

	for i := 0; i < len(nums); i++ {
		max += nums[i]

		if res < max {
			res = max
		}
		if max < 0 {
			max = 0
		}
	}

	return res
}
