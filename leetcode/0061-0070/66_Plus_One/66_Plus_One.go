package main

import (
	"fmt"
)

func main() {
	test([]int{0}, []int{1})
	test([]int{9}, []int{1, 0})
	test([]int{1, 9}, []int{2, 0})
	test([]int{1, 2, 3}, []int{1, 2, 4})
}

func test(dight []int, target []int) {
	temp := plusOne(dight)
	for i := 0; i < len(temp); i++ {
		if temp[i] != target[i] {
			fmt.Println("================", dight, target)
			break
		}
	}
}

func plusOne(digits []int) []int {
	digits = append([]int{0}, digits...)
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		temp := 0
		if i == len(digits)-1 {
			temp = digits[i] + 1 + carry
		} else {
			temp = digits[i] + carry
		}
		digits[i] = temp % 10
		carry = temp / 10
	}
	if digits[0] == 0 {
		digits = digits[1:]
	}
	return digits
}
