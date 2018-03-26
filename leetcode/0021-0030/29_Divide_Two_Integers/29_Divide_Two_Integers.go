package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0
	b := 1
	fmt.Println(divide(a, b))
	a = -2147483648
	b = 2
	fmt.Println(divide(a, b))
	fmt.Print(2 >> 1)
	fmt.Print(2 << 0)
	fmt.Print(2 << 1)
}

func divide(dividend int, divisor int) int {
	sign := 1
	if dividend < 0 {
		dividend = -dividend
		sign = -sign
	}
	if divisor < 0 {
		divisor = -divisor
		sign = -sign
	}

	var count, i uint64
	for dividend >= divisor {
		if dividend < divisor<<i {
			dividend -= divisor << (i - 1)
			count += (2 << (i - 1)) >> 1
			i = 0
		} else if dividend == divisor<<i {
			count += (2 << i) >> 1
			break
		} else {
			i++
		}
	}

	result := int(count) * sign
	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}
	return result
}
