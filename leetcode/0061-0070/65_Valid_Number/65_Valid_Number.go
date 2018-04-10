package main

import (
	"fmt"
	"regexp"
)

func main() {
	test(1, "123", true)
	test(2, " 123 ", true)
	test(3, "0", true)
	test(4, "0123", true) //Cannot agree
	test(5, "00", true)   //Cannot agree
	test(6, "-10", true)
	test(7, "-0", true)
	test(8, "123.5", true)
	test(9, "123.000000", true)
	test(10, "-500.777", true)
	test(11, "0.0000001", true)
	test(12, "0.00000", true)
	test(13, "0.", true)   //Cannot be more disagree!!!
	test(14, "00.5", true) //Strongly cannot agree
	test(15, "123e1", true)
	test(16, "1.23e10", true)
	test(17, "0.5e-10", true)
	test(18, "1.0e4.5", false)
	test(19, "0.5e04", true)
	test(20, "12 3", false)
	test(21, "1a3", false)
	test(22, "", false)
	test(23, "     ", false)
	test(24, ".1", true)         //Ok, if you say so
	test(25, " 005047e+6", true) //Damn = =|||
	test(26, ".", false)
	test(27, "2e0", true) //Really?!
	test(28, "+.8", true)
}

func test(num int, s string, res bool) {
	if res == isNumber(s) {
		fmt.Println(num, s, res, "ok")
	} else {
		fmt.Println("+++++++++++++++++++", num, s, res, "no!!!!!")
	}
}
func isNumber(s string) bool {
	reg := regexp.MustCompile(`([\ ])*([\+|-]?)([0-9]+[\.]?[0-9]*|[\.][0-9]+)(e([\+|-]?)([0-9]+))?([\ ])*`)
	fmt.Println(reg.FindString(s))
	res := reg.FindString(s)
	if len(res) == 0 {
		return false
	}
	return s == res
}
