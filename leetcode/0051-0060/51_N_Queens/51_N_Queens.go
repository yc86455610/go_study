package main

import "fmt"

func main() {
	a := 4
	fmt.Println(solveNQueens(a))

	b := [][]byte{
		{'Q', '.', '.', '.'},
		{'.', '.', 'Q', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	fmt.Println(isValid(b, 1, 1))
}

func solveNQueens(n int) [][]string {
	// init
	bSlice := [][]byte{}
	for i := 0; i < n; i++ {
		t := []byte{}
		for j := 0; j < n; j++ {
			t = append(t, '.')
		}
		bSlice = append(bSlice, t)
	}
	// fmt.Println("bslice", bSlice)
	res := [][]string{}

	solve(bSlice, 0, &res)
	// result
	return res
}

func solve(a [][]byte, row int, res *[][]string) {
	// fmt.Println(a)
	if row == len(a) {
		addByteSliceToResult(a, res)
		return
	}
	for i := 0; i < len(a); i++ {
		if isValid(a, row, i) {
			a[row][i] = 'Q'
			solve(a, row+1, res)
			a[row][i] = '.'
		}
	}
}

func isValid(a [][]byte, row, col int) bool {
	// check row/col
	for i := 0; i < len(a); i++ {
		if i != col && a[row][i] == 'Q' {
			return false
		}
		if i != row && a[i][col] == 'Q' {
			return false
		}
	}

	// check 45 degree
	for i, j := row-1, col+1; i >= 0 && j < len(a); {
		if a[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	for i, j := row+1, col-1; i < len(a) && j >= 0; {
		if a[i][j] == 'Q' {
			return false
		}
		i++
		j--
	}

	// check 135 degree
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if a[i][j] == 'Q' {
			return false
		}
		i--
		j--

	}
	for i, j := row+1, col+1; i < len(a) && j < len(a); {
		if a[i][j] == 'Q' {
			return false
		}
		i++
		j++
	}
	return true
}

func addByteSliceToResult(a [][]byte, res *[][]string) {
	r := []string{}
	for i := 0; i < len(a); i++ {
		r = append(r, string(a[i]))
	}
	*res = append(*res, r)
}
