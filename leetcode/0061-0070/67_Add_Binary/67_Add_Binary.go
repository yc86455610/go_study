package main

import (
	"fmt"
)

func main() {
	test("0", "0", "0")
	test("1", "0", "1")
	test("0", "1", "1")
	test("1", "1", "10")
	test("10", "1", "11")
	test("1", "10", "11")
	test("11", "1", "100")
	test("1", "11", "100")
	test("11", "11", "110")
	test("101111", "10", "110001")
}

func test(a, b, c string) {
	if c != addBinary(a, b) {
		fmt.Println("================", a, b, c)
	}
}

func addBinary(a string, b string) string {
	maxLen := 0
	if len(a) > len(b) {
		maxLen = len(a) + 1
	} else {
		maxLen = len(b) + 1
	}
	byteA := []byte(a)
	byteB := []byte(b)
	for i := 0; i < maxLen-len(a); i++ {
		byteA = append([]byte{'0'}, byteA...)
	}
	for i := 0; i < maxLen-len(b); i++ {
		byteB = append([]byte{'0'}, byteB...)
	}
	var carry byte
	carry = '0'
	for i := maxLen - 1; i >= 0; i-- {
		temp := byteA[i] - '0' + byteB[i] - '0' + carry - '0'
		byteA[i] = temp%2 + '0'
		carry = (temp / 2) + '0'
	}
	if byteA[0] == '0' {
		byteA = byteA[1:]
	}
	return string(byteA)
}
