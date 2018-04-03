package main

import "fmt"

func main() {
	a := 2.5
	fmt.Println(myPow(a, 16))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	res := 0.0
	if n%2 == 0 {
		temp := myPow(x, n/2)
		fmt.Println(temp)
		res = temp * temp
	} else {
		temp := myPow(x, n/2)
		fmt.Println(temp)
		res = x * temp * temp
	}
	return res
}
