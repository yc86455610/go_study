package main

import (
	"fmt"
	"strings"
)

func main() {
	test("/", "/")
	test("/../", "/")
	test("/a/", "/a")
	test("/a/.", "/a")
	test("/a/./b", "/a/b")
	test("/a/./b/..", "/a")
	test("/a//.//b/..", "/a")
}

func test(path, SPath string) {
	if SPath != simplifyPath(path) {
		fmt.Println(path, simplifyPath(path), SPath)
	}
}

func simplifyPath(path string) string {
	s := strings.Split(path, "/")
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == "" || s[i] == "." {
			continue
		} else if s[i] == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	res := []byte{}
	for i := 0; i < len(stack); i++ {
		res = append(res, '/')
		res = append(res, []byte(stack[i])...)
	}
	if len(stack) == 0 {
		res = append(res, '/')
	}
	return string(res)
}
