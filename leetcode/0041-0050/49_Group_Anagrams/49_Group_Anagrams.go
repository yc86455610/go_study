package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(a))
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	m := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		temp := sortString(strs[i])
		m[temp] = append(m[temp], strs[i])
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
