package main

import "fmt"

func main() {
	a := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(a))
}

func canJump(nums []int) bool {
	maxRearch := 0
	for i := 0; i < len(nums); i++ {
		if maxRearch < i {
			return false
		}
		if maxRearch < nums[i]+i {
			maxRearch = nums[i] + i
		}
	}

	return true
}
