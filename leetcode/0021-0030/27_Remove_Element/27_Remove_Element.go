package main

import "fmt"

func main() {
	a := []int{3, 2, 3, 2, 3}
	b := 3
	fmt.Println(removeElement(a, b))
	a = []int{1}
	b = 1
	fmt.Println(removeElement(a, b))
	a = []int{2, 2, 3}
	b = 2
	fmt.Println(removeElement(a, b))
}

func removeElement(nums []int, val int) int {
	tail := len(nums)
	for i := 0; i < tail; i++ {
		for tail != 0 && nums[tail-1] == val {
			tail--
		}
		if tail == 0 || i >= tail {
			break
		}
		if nums[i] == val {
			nums[i], nums[tail-1] = nums[tail-1], nums[i]
			tail--
		}
	}
	return tail
}
