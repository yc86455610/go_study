package main

import "fmt"

func main() {
	for i := 1; i <= 24; i++ {
		fmt.Println(getPermutation(4, i))
	}
}

func getPermutation(n int, k int) string {
	data := make([]int, n+1)
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = byte(i + 1 + '0')
	}
	data[0] = 1
	for i := 1; i < len(data); i++ {
		data[i] = i * data[i-1]
	}
	// fmt.Println(data, bytes)

	res := make([]byte, n)

	get(n, k-1, &res, 0, data, &bytes)

	return string(res)
}

// n: n种元素排列，k:以0为开始的第k个排序序列，res:返回的字符串，location:从左开始第location位，
// data: 切片，第i位表示i种元素的总排序数，bytes：切片，保存还剩的字符串集合
func get(n int, k int, res *[]byte, location int, data []int, bytes *[]byte) {
	if location >= n {
		return
	}
	b := k / data[n-1-location]
	c := k % data[n-1-location]
	// println(b, c)
	(*res)[location] = (*bytes)[b]
	*bytes = append((*bytes)[:b], (*bytes)[b+1:]...)
	// fmt.Println(*bytes)
	get(n, c, res, location+1, data, bytes)
}
