package main

import "fmt"

func main() {
	a := [][]byte{}
	a = append(a, []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'})
	a = append(a, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'})
	a = append(a, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'})
	a = append(a, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'})
	a = append(a, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'})
	a = append(a, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'})
	a = append(a, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'})
	a = append(a, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'})
	a = append(a, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'})
	solveSudoku(a)
	display(a)
}

func display(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%s ", string(board[i][j]))
		}
		fmt.Println()
	}
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' && board[row][i] == num {
			return false
		}
		if board[i][col] != '.' && board[i][col] == num {
			return false
		}
		if board[(row/3)*3+i/3][(col/3)*3+i%3] != '.' && board[(row/3)*3+i/3][(col/3)*3+i%3] == num {
			return false
		}
	}
	return true
}
func solve(board *[][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*board)[i][j] == '.' {
				for k := '1'; k <= '9'; k++ {
					if isValid(*board, i, j, byte(k)) {
						(*board)[i][j] = byte(k)
						if solve(board) {
							return true
						}
						(*board)[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	solve(&board)
}
