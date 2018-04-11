package main

import (
	"fmt"
)

func main() {
	test([][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}, 3, true)
	test([][]int{}, 0, false)
	test([][]int{[]int{}}, 1, false)
	test([][]int{[]int{1}}, 1, true)
}

func test(matrix [][]int, target int, b bool) {
	if b != searchMatrix(matrix, target) {
		fmt.Println(matrix, target)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	row, res := findRow(matrix, target)
	if res {
		return res
	}
	res = findCol(matrix, target, row)

	return res
}

func findRow(matrix [][]int, target int) (int, bool) {
	m := len(matrix)
	if target >= matrix[m-1][0] {
		return m - 1, false
	}
	left := 0
	right := m - 1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] < target {
			left = mid + 1
		} else if matrix[mid][0] > target {
			right = mid - 1
		} else {
			return mid, true
		}
	}
	return right, false
}

func findCol(matrix [][]int, target int, row int) bool {
	n := len(matrix[0])
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if matrix[row][mid] < target {
			left = mid + 1
		} else if matrix[row][mid] > target {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
