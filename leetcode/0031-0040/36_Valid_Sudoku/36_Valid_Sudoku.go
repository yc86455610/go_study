package main

import "fmt"

//[[],[],[],[],[],[],[],[],[]]

func main() {
	a := [][]byte{}
	a = append(a, []byte{'.', '8', '7', '6', '5', '4', '3', '2', '1'})
	a = append(a, []byte{'2', '.', '.', '.', '.', '.', '.', '.', '.'})
	a = append(a, []byte{'3', '.', '.', '.', '.', '.', '.', '.', '.'})
	a = append(a, []byte{'4', '.', '.', '.', '.', '.', '.', '.', '.'})
	a = append(a, []byte{'5', '.', '.', '.', '.', '.', '.', '.', '.'})
	a = append(a, []byte{'6', '.', '.', '.', '.', '.', '.', '.', '.'})
	a = append(a, []byte{'7', '.', '.', '.', '.', '.', '.', '.', '.'})
	a = append(a, []byte{'8', '.', '.', '.', '.', '.', '.', '.', '.'})
	a = append(a, []byte{'9', '.', '.', '.', '.', '.', '.', '.', '.'})
	fmt.Println(isValidSudoku(a))
}

func isValidSudoku(board [][]byte) bool {
	return isValidRow(board) && isValidCol(board) && isValidArea(board)
}

func isValidRow(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		set := map[byte]int{}
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				set[board[i][j]]++
				if set[board[i][j]] != 1 {
					return false

				}
			}
		}
	}
	return true
}

func isValidCol(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		set := map[byte]int{}
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				set[board[j][i]]++
				if set[board[j][i]] != 1 {
					return false
				}
			}
		}
	}
	return true
}

func isValidArea(board [][]byte) bool {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			set := map[byte]int{}
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[k+i][l+j] != '.' {
						set[board[k+i][l+j]]++
						if set[board[k+i][l+j]] != 1 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
