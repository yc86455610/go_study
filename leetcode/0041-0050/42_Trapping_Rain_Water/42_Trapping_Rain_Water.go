package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{}
	fmt.Println(trap(a))
}

func trap(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1
	minH := math.MaxInt16
	for left < right {
		if height[left] < height[right] {
			minH = height[left]
			for left < right && height[left] <= minH {
				res += minH - height[left]
				left++
			}
		} else {
			minH = height[right]
			for left < right && height[right] <= minH {
				res += minH - height[right]
				right--
			}
		}
	}
	return res
}
