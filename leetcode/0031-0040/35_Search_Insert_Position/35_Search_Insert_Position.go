package main

import "fmt"

func main() {
	a := []int{1, 3, 4, 6}
	fmt.Println(searchInsert(a, 2))
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}
