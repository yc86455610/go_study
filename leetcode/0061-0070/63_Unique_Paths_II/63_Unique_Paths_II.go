package main

import "fmt"

func main() {
	fmt.Println()
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}
	left := 0
	up := 0
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == -1 {
				continue
			}
			if i == 0 && j == 0 {
				if obstacleGrid[i][j] == 0 {
					obstacleGrid[i][j] = 1
					continue
				}
			}
			if i == 0 || obstacleGrid[i-1][j] == -1 {
				up = 0
			} else {
				up = obstacleGrid[i-1][j]
			}
			if j == 0 || obstacleGrid[i][j-1] == -1 {
				left = 0
			} else {
				left = obstacleGrid[i][j-1]
			}
			obstacleGrid[i][j] = up + left
		}
	}
	res := obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
	if res == -1 {
		return 0
	} else {
		return res
	}
}
