package main

import (
	"fmt"
)

func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	var rNum int
	for x > 0 {
		rNum = rNum*10 + x%10
		x = x / 10
	}
	return rNum
}

func main() {
	a := -123
	fmt.Println(reverse(a))
	b := 123
	fmt.Println(reverse(b))
}
