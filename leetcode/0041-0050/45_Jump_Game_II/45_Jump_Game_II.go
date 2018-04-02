package main

import "fmt"

func main() {
	a := []int{0}
	fmt.Println(jump(a))
}

func jump(nums []int) int {
	step := 0
	lastRearch := 0
	maxRearch := 0
	for i := 0; i < len(nums); i++ {
		if maxRearch < i {
			return -1
		}
		if lastRearch < i {
			step++
			lastRearch = maxRearch
		}
		if maxRearch < nums[i]+i {
			maxRearch = nums[i] + i
		}
	}
	return step
}
