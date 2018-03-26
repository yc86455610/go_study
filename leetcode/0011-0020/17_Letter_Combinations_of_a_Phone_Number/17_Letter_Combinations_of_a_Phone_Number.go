package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := "1243"
	fmt.Println(letterCombinations(a))
}

func letterCombinations(digits string) []string {
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	rString := []string{}
	for i := 0; i < len(digits); i++ {
		tmpString := []string{}

		if len(rString) == 0 {
			for j := 0; j < len(m[digits[i]-'0']); j++ {
				tmpString = append(tmpString, string(m[digits[i]-'0'][j]))
			}
		}

		for j := 0; j < len(rString); j++ {
			for k := 0; k < len(m[digits[i]-'0']); k++ {
				b := bytes.Buffer{}
				b.WriteString(rString[j])
				b.WriteByte(m[digits[i]-'0'][k])
				tmpString = append(tmpString, b.String())
			}
		}
		rString = tmpString
	}
	return rString
}
