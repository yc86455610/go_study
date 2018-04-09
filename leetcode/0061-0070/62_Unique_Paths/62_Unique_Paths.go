package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 4))
}

func uniquePaths(m int, n int) int {
	myMap := make([][]int, 0)
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		myMap = append(myMap, temp)
	}
	fmt.Println(myMap)
	myMap[0][0] = 1
	left := 0
	up := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				up = 0
			} else {
				up = myMap[i-1][j]
			}
			if j == 0 {
				left = 0
			} else {
				left = myMap[i][j-1]
			}
			myMap[i][j] = up + left
		}
	}
	fmt.Println(myMap)
	return myMap[m-1][n-1]
}

// func uniquePaths(m int, n int) int {
// 	count := 0
// 	dfs(m, n, 0, 0, &count)
// 	return count
// }

// func dfs(m int, n int, row int, col int, count *int) {
// 	if m == row+1 && n == col+1 {
// 		*count++
// 		return
// 	}
// 	if row < m-1 {
// 		dfs(m, n, row+1, col, count)
// 	}
// 	if col < n-1 {
// 		dfs(m, n, row, col+1, count)
// 	}
// }
