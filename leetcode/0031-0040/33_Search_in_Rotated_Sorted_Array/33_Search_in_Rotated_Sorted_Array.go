package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 5, 8, -1, 0}
	fmt.Println(search(a, 8))
	a = []int{1, 3}
	fmt.Println(search(a, 2))
	a = []int{1, 3}
	fmt.Println(search(a, 0))
	a = []int{3, 1}
	fmt.Println(search(a, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return rotatedBinarySearch(nums, 0, len(nums)-1, target)
}

func rotatedBinarySearch(nums []int, left, right, target int) int {
	index := (left + right) / 2
	if len(nums[left:right+1]) == 1 {
		if target == nums[index] {
			return index
		}
		return -1
	}
	if nums[index] > nums[left] {
		if result := binarySearch(nums, left, index, target); result != -1 {
			return result
		}
	} else {
		if result := rotatedBinarySearch(nums, left, index, target); result != -1 {
			return result
		}
	}
	if nums[index] < nums[right] {
		if result := binarySearch(nums, index, right, target); result != -1 {
			return result
		}
	} else {
		if result := rotatedBinarySearch(nums, index+1, right, target); result != -1 {
			return result
		}
	}
	return -1
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		index := (left + right) / 2
		if target == nums[index] {
			return index
		} else if target > nums[index] {
			left = index + 1
		} else if target < nums[index] {
			right = index - 1
		}
	}
	return -1
}
