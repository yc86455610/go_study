package main

import (
	"fmt"
)

func main() {
	var text string
	var nRows int
	var r string
	text = "PAYPALISHIRING"
	nRows = 3
	r = convert(text, nRows)
	fmt.Println(r)
	text = "PAYPALISHIRING"
	nRows = 1
	r = convert(text, nRows)
	fmt.Println(r)
	text = ""
	nRows = 1
	r = convert(text, nRows)
	fmt.Println(r)
}
func convert(s string, numRows int) string {
	var buffer [][]byte
	var rByte []byte
	for i := 0; i < numRows; i++ {
		buffer = append(buffer, nil)
	}
	var i, row int
	var isDown = true
	for i = 0; i < len(s); i++ {
		if isDown {
			if row < numRows-1 {
				buffer[row] = append(buffer[row], s[i])
				row++
			} else if row == numRows-1 {
				buffer[row] = append(buffer[row], s[i])
				row--
				isDown = false
			} else {
				row--
				isDown = false
			}
		} else {
			if row > 0 {
				buffer[row] = append(buffer[row], s[i])
				row--
			} else if row == 0 {
				buffer[row] = append(buffer[row], s[i])
				row++
				isDown = true
			} else {
				row++
				isDown = true
				i--
			}
		}
	}
	for i := 0; i < numRows; i++ {
		rByte = append(rByte, buffer[i]...)
	}
	return string(rByte)
}
