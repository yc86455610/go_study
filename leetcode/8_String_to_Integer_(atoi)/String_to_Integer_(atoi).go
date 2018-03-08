package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	var rNum int
	sign := 1
	signCount := 0
	preSpace := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			preSpace++
		} else {
			break
		}
	}
	for i := preSpace; i < len(str); i++ {
		if str[i] == '-' {
			sign = -1
			signCount++
		} else if str[i] == '+' {
			sign = 1
			signCount++
		} else if str[i] >= '0' && str[i] <= '9' {
			rNum = rNum*10 + (int(str[i]) - int('0'))
			if rNum > math.MaxInt32 {
				if sign == 1 {
					return math.MaxInt32
				} else if sign == -1 {
					return math.MinInt32
				}
			}
		} else {
			break
		}
	}
	if signCount > 1 {
		return 0
	}
	return rNum * sign

}

func main() {
	s := ""
	fmt.Println(myAtoi(s))
	s = "+-2"
	fmt.Println(myAtoi(s))
	s = "   +0 123"
	fmt.Println(myAtoi(s))
	s = "9223372036854775809"
	fmt.Println(myAtoi(s))
	s = "-123"
	fmt.Println(myAtoi(s))
}
