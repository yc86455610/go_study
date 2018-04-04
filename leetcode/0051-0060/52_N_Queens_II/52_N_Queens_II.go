package main

import "fmt"

func main() {
	fmt.Println(totalNQueens(1))
}

func totalNQueens(n int) int {
	res := 0
	queen := make([]int, n)
	solve(queen, 0, &res)
	return res
}

func solve(queen []int, step int, res *int) {
	if step == len(queen) {
		*res++
		return
	}
	for i := 0; i < len(queen); i++ {
		if isValid(queen, step, i) {
			queen[step] = i
			solve(queen, step+1, res)
		}
	}
}

func isValid(queen []int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if i-queen[i] == row-col || i+queen[i] == row+col || queen[i] == col {
			return false
		}
	}
	return true
}
