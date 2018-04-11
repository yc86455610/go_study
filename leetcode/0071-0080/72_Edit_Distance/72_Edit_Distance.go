package main

import (
	"fmt"
	"math"
)

func main() {
	test("", "", 0)
	test("1", "", 1)
	test("1", "12", 1)
	test("123", "21", 2)
}

func test(word1, word2 string, res int) {
	if res != minDistance(word1, word2) {
		fmt.Println(word1, word2, res)
	}
}

func minDistance(word1 string, word2 string) int {
	m := [][]int{}
	for i := 0; i < len(word1)+1; i++ {
		temp := make([]int, len(word2)+1)
		m = append(m, temp)
	}
	for i := 0; i < len(word1)+1; i++ {
		m[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		m[0][i] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if m[i][j] != 0 {
				continue
			}
			temp := math.MaxInt64
			if m[i-1][j]+1 < temp {
				temp = m[i-1][j] + 1
			}
			if m[i][j-1]+1 < temp {
				temp = m[i][j-1] + 1
			}
			if word1[i-1] == word2[j-1] {
				if m[i-1][j-1] < temp {
					temp = m[i-1][j-1]
				}
			} else {
				if m[i-1][j-1]+1 < temp {
					temp = m[i-1][j-1] + 1
				}
			}
			m[i][j] = temp
		}
	}

	return m[len(word1)][len(word2)]
}
