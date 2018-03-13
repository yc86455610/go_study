package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := "{}}{{}()"
	fmt.Println(isValid(a))
	a = "()"
	fmt.Println(isValid(a))
}

func isValid(s string) bool {
	l := list.New()

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			l.PushBack(s[i])
		} else {
			if l.Len() != 0 {
				tmp := l.Back()
				l.Remove(tmp)
				v := int(tmp.Value.(byte))
				if (v != '{' && s[i] == '}') || (v != '[' && s[i] == ']') || (v != '(' && s[i] == ')') {
					return false
				}
			} else {
				return false
			}
		}
	}
	if l.Len() != 0 {
		return false
	}
	return true
}
