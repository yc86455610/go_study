package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 4, 6}
	fmt.Println(searchRange(a, 4))
}

func searchRange(nums []int, target int) []int {
	left := binarySearchLeft(nums, 0, len(nums)-1, target)
	right := binarySearchRight(nums, 0, len(nums)-1, target)
	if left <= right {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func binarySearchLeft(nums []int, left, right, target int) int {
	if left > right {
		return left
	}
	mid := (left + right) / 2
	if nums[mid] < target {
		return binarySearchLeft(nums, mid+1, right, target)
	}
	return binarySearchLeft(nums, left, mid-1, target)
}

func binarySearchRight(nums []int, left, right, target int) int {
	if left > right {
		return right
	}
	mid := (left + right) / 2
	if nums[mid] > target {
		return binarySearchRight(nums, left, mid-1, target)
	}
	return binarySearchRight(nums, mid+1, right, target)
}
