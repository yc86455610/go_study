package main

import (
	"fmt"
)

func main() {
	test([][]byte{[]byte{'a'}}, "ab", false)
}

func test(b [][]byte, word string, bl bool) {
	if bl != exist(b, word) {
		fmt.Println(b, word, bl)
	}
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	m := len(board)
	n := len(board[0])
	t := [][]bool{}
	for i := 0; i < m; i++ {
		temp := make([]bool, n)
		t = append(t, temp)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] == board[i][j] {
				t[i][j] = true
				if dfs(board, i, j, word, 1, &t) {
					return true
				}
				t[i][j] = false
			}
		}
	}
	return false
}

func dfs(board [][]byte, row int, col int, word string, index int, t *[][]bool) bool {
	if len(word) == index {
		return true
	}

	if row != 0 && board[row-1][col] == word[index] && (*t)[row-1][col] == false {
		(*t)[row-1][col] = true
		if dfs(board, row-1, col, word, index+1, t) {
			return true
		}
		(*t)[row-1][col] = false
	}
	if col != 0 && board[row][col-1] == word[index] && (*t)[row][col-1] == false {
		(*t)[row][col-1] = true
		if dfs(board, row, col-1, word, index+1, t) {
			return true
		}
		(*t)[row][col-1] = false
	}
	if row != len(board)-1 && board[row+1][col] == word[index] && (*t)[row+1][col] == false {
		(*t)[row+1][col] = true
		if dfs(board, row+1, col, word, index+1, t) {
			return true
		}
		(*t)[row+1][col] = false
	}
	if col != len(board[0])-1 && board[row][col+1] == word[index] && (*t)[row][col+1] == false {
		(*t)[row][col+1] = true
		if dfs(board, row, col+1, word, index+1, t) {
			return true
		}
		(*t)[row][col+1] = false
	}
	return false
}
