package main

import (
	"fmt"
)

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(a)
	rotate(a)
	fmt.Println(a)
}

func rotate(matrix [][]int) {
	// 上下翻
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-1-i] = matrix[len(matrix)-1-i], matrix[i]
	}

	// 左上至右下对角线翻
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
