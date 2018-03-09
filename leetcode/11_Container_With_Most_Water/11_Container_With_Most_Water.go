package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 1}
	fmt.Println(maxArea(a))
}

func maxArea(height []int) int {
	// 初始化,height开头插入0，方便计算
	height = append([]int{0}, height[:]...)

	// 从两边向中心缩减，每次比较找到更大的容量，每次移动较短高度的坐标。
	lIndex := 0
	rIndex := len(height) - 1
	rMaxArea := 0
	for lIndex < rIndex {
		if (rIndex-lIndex)*int(math.Min(float64(height[lIndex]), float64(height[rIndex]))) > rMaxArea {
			rMaxArea = (rIndex - lIndex) * int(math.Min(float64(height[lIndex]), float64(height[rIndex])))
		}
		if height[lIndex] < height[rIndex] {
			lIndex++
		} else {
			rIndex--
		}
	}
	return rMaxArea
}
