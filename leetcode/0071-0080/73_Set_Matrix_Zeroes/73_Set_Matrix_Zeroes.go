package main

func main() {
	setZeroes([][]int{[]int{0}})
}

func setZeroes(matrix [][]int) {
	var row, col bool
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			row = true
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			col = true
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < m; i++ {
		if row {
			matrix[i][0] = 0
		}
	}
	for i := 0; i < n; i++ {
		if col {
			matrix[0][i] = 0
		}
	}
}
