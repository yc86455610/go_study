package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(countAndSay(a))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		tmp := []byte{}
		index := 0
		for j := 0; j < len(str); j++ {
			if j == 0 {
				tmp = append(tmp, '1')
				tmp = append(tmp, str[0])
				continue
			}
			if str[j] == str[j-1] {
				tmp[index]++
			} else {
				index += 2
				tmp = append(tmp, '1')
				tmp = append(tmp, str[j])
			}
		}
		str = string(tmp)
	}
	return str
}
