package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	sum := []int{}

	for i, number := range nums {
		for j, number1 := range nums[i+1:] {
			if number+number1 == target {
				sum = append(sum, i, (i + j + 1))
			}
		}
	}
	return sum
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := twoSum(nums, target)
	fmt.Println(sum)
}
