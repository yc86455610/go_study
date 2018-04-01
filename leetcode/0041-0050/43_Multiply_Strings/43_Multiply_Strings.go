package main

import "fmt"

func main() {
	a := "192304"
	b := "92304"
	fmt.Println(multiply(a, b))
}
func multiply(num1 string, num2 string) string {
	if len(num1) == 1 && num1[0] == '0' || len(num2) == 1 && num2[0] == '0' {
		return "0"
	}
	a := []byte(num1)
	b := []byte(num2)
	reverse(a)
	reverse(b)
	res := []byte{}

	for i := 0; i < len(b); i++ {
		res = add(res, mul(a, b[i], i+1))
	}
	reverse(res)
	return string(res)
}

func mul(a []byte, b byte, digit int) []byte {
	res := []byte{}
	for i := 0; i < digit-1; i++ {
		res = append(res, '0')
	}
	carry := 0
	for i := 0; i < len(a); i++ {
		res = append(res, byte(int((a[i]-'0')*(b-'0'))+carry)%10+'0')
		carry = (int((a[i]-'0')*(b-'0')) + carry) / 10
	}
	if carry != 0 {
		res = append(res, byte(carry)+'0')
	}
	return res
}

func add(a []byte, b []byte) []byte {
	res := []byte{}
	carry := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		tmpA := 0
		tmpB := 0
		if i < len(a) {
			tmpA = int(a[i]) - '0'
		}
		if i < len(b) {
			tmpB = int(b[i]) - '0'
		}
		res = append(res, byte((tmpA+tmpB+carry)%10+'0'))
		carry = (tmpA + tmpB + carry) / 10
	}
	if carry != 0 {
		res = append(res, byte(carry)+'0')
	}
	return res
}

// reverse reverses a slice of ints in place.
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
