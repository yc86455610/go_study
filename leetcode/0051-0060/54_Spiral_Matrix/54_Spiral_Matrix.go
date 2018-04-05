package main

import "fmt"

func main() {
	fmt.Println()
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	rowBegin := 0
	rowEnd := len(matrix) - 1
	colBegin := 0
	colEnd := len(matrix[0]) - 1
	fmt.Print(rowBegin, rowEnd, colBegin, colEnd)

	for n := len(matrix[0]) * len(matrix); n > 0; {
		for i := colBegin; i <= colEnd; i++ {
			res = append(res, matrix[rowBegin][i])
			n--
		}
		rowBegin++
		for i := rowBegin; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
			n--
		}
		colEnd--
		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				res = append(res, matrix[rowEnd][i])
				n--
			}
		}
		rowEnd--
		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				res = append(res, matrix[i][colBegin])
				n--
			}
		}
		colBegin++
	}

	return res
}
