package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := 1234
	fmt.Println(intToRoman(a))
}

func intToRoman(num int) string {
	// 建表，得到数字对应的罗马字符
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	// 把数字转换成字符连接起来
	var buf bytes.Buffer
	buf.WriteString(M[num/1000])
	buf.WriteString(C[(num%1000)/100])
	buf.WriteString(X[(num%100)/10])
	buf.WriteString(I[num%10])

	return buf.String()
}
