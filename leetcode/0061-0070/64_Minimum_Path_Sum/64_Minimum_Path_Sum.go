package main

import "math"

func main() {

}

func minPathSum(grid [][]int) int {
	res := 0
	up := 0
	left := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				up = math.MaxInt64
			} else {
				up = grid[i-1][j]
			}
			if j == 0 {
				left = math.MaxInt64
			} else {
				left = grid[i][j-1]
			}
			if up > left {
				grid[i][j] += left
			} else if up < left {
				grid[i][j] += up
			} else {
				if up != math.MaxInt64 {
					grid[i][j] += up
				}
			}
		}
	}
	res = grid[len(grid)-1][len(grid[0])-1]
	return res
}
