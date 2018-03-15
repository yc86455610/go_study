package main

import "fmt"

func main() {
	a := []int{1, 1, 2}
	fmt.Println(removeDuplicates(a))
	fmt.Println(a)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
