package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{2, 3, 1, 3, 3}
	nextPermutation(a)
	fmt.Println(a)
	a = []int{1, 3, 2}
	nextPermutation(a)
	fmt.Println(a)
}

func nextPermutation(nums []int) {
	i := len(nums) - 1
	min := math.MaxInt32
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			k := i
			for j := i; j < len(nums); j++ {
				if nums[i-1] < nums[j] && nums[j] <= min {
					k = j
					min = nums[j]
				}
			}
			nums[k], nums[i-1] = nums[i-1], nums[k]
			break
		}
	}
	for j := 0; j <= (len(nums)-1-i)/2; j++ {
		nums[i+j], nums[len(nums)-1-j] = nums[len(nums)-1-j], nums[i+j]
	}
}
