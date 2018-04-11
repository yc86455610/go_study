package main

import "fmt"

func main() {
	test(1, 1)
	test(2, 1)
	test(3, 1)
	test(4, 2)
	test(5, 2)
}

func test(x int, y int) {
	if y != mySqrt(x) {
		fmt.Printf("wrong answer. x: %d, mySqrt:%d, expacted: %d\n", x, mySqrt(x), y)
	}
}

func mySqrt(x int) int {
	res := 1
	for {
		if res*res > x {
			break
		} else {
			res++
		}
	}
	return res - 1
}
