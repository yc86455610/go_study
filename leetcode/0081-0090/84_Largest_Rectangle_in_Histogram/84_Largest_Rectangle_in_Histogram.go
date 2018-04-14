package main

import (
	"fmt"
)

func main() {
	// test([]int{2, 1, 5, 6, 2, 3}, 10)
	test([]int{2, 1, 2}, 3)
}

func test(heights []int, r int) {
	if r != largestRectangleArea(heights) {
		fmt.Println(heights, r)
	}
}

func largestRectangleArea(heights []int) int {
	res := 0
	stack := []int{}
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 || stack[len(stack)-1] <= heights[i] {
			stack = append(stack, heights[i])
		} else {
			count := 0
			for len(stack) != 0 && stack[len(stack)-1] >= heights[i] {
				temp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				count++
				if res < temp*count {
					res = temp * count
				}
			}
			for j := 0; j < count+1; j++ {
				stack = append(stack, heights[i])
			}
		}
	}
	return res
}
