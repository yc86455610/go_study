package main

import "fmt"

func main() {
	a := 2
	fmt.Println(generateMatrix(a))
}

func generateMatrix(n int) [][]int {
	res := [][]int{}
	for i := 0; i < n; i++ {
		temp := make([]int, n)
		res = append(res, temp)
	}
	fmt.Println(res)
	rowBegin := 0
	rowEnd := n - 1
	colBegin := 0
	colEnd := n - 1
	for num := 1; num <= n*n; {
		for i := colBegin; i <= colEnd; i++ {
			res[rowBegin][i] = num
			num++
		}
		rowBegin++
		fmt.Println(res)
		for i := rowBegin; i <= rowEnd; i++ {
			res[i][colEnd] = num
			num++
		}
		colEnd--
		fmt.Println(res)
		if colBegin <= colEnd {
			for i := colEnd; i >= colBegin; i-- {
				res[rowEnd][i] = num
				num++
			}
		}
		rowEnd--
		fmt.Println(res)
		if rowBegin <= rowEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				res[i][colBegin] = num
				num++
			}
		}
		colBegin++
		fmt.Println(res)
	}

	return res
}
