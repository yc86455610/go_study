package main

import "fmt"

func main() {
	test([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{"This    is    an", "example  of text", "justification.  "})
	test([]string{"a", "b", "c"}, 1, []string{"a", "b", "c"})
	test([]string{"a", "b", "c"}, 2, []string{"a ", "b ", "c "})
	test([]string{"a", "b", "c"}, 3, []string{"a b", "c  "})
	test([]string{"a", "b", "c", "d"}, 4, []string{"a  b", "c d "})
	test([]string{"a", "b", "c"}, 5, []string{"a b c"})
}

func test(words []string, maxWidth int, res []string) {
	r := fullJustify(words, maxWidth)
	for i := 0; i < len(res); i++ {
		if res[i] != r[i] {
			fmt.Println(words, maxWidth, r, res)
			fmt.Println(r[i], res[i])
		}
	}
}

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}

	count := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		count[i] = len(words[i])
	}
	group := [][]int{}
	temp := []int{}
	groupWidth := []int{}
	groupWidthCount := 0
	remainWidth := maxWidth
	for i := 0; i < len(count); i++ {
		if remainWidth > count[i] {
			temp = append(temp, i)
			groupWidthCount = groupWidthCount + count[i]
			remainWidth = remainWidth - count[i] - 1
			if i == len(count)-1 {
				group = append(group, temp)
				groupWidth = append(groupWidth, groupWidthCount)
			}
		} else if remainWidth == count[i] {
			temp = append(temp, i)
			groupWidthCount = groupWidthCount + count[i]
			remainWidth = remainWidth - count[i]
			if i == len(count)-1 {
				group = append(group, temp)
				groupWidth = append(groupWidth, groupWidthCount)
			}
		} else {
			group = append(group, temp)
			groupWidth = append(groupWidth, groupWidthCount)
			groupWidthCount = 0
			remainWidth = maxWidth
			temp = []int{}
			i--
		}
	}

	bArr := [][]byte{}

	for i := 0; i < len(group)-1; i++ {
		btemp := []byte{}
		lastSpaceCount := maxWidth - groupWidth[i]
		remainSpaceCount := lastSpaceCount
		for j := 0; j < len(group[i]); j++ {
			spaceCount := 0
			if len(group[i])-j-1 != 0 {
				spaceCount = remainSpaceCount / (len(group[i]) - j - 1)
				if (remainSpaceCount % (len(group[i]) - j - 1)) != 0 {
					spaceCount++
				}
			} else {
				spaceCount = remainSpaceCount
			}
			remainSpaceCount = remainSpaceCount - spaceCount
			addByte(&btemp, words[group[i][j]], spaceCount)
		}
		bArr = append(bArr, btemp)
	}
	btemp := []byte{}
	remainSpaceCount := maxWidth - groupWidth[len(group)-1]
	for j := 0; j < len(group[len(group)-1]); j++ {
		spaceCount := 1
		if len(group[len(group)-1])-j-1 == 0 {
			spaceCount = remainSpaceCount
		}
		remainSpaceCount = remainSpaceCount - spaceCount
		addByte(&btemp, words[group[len(group)-1][j]], spaceCount)
	}
	bArr = append(bArr, btemp)
	for i := 0; i < len(bArr); i++ {
		res = append(res, string(bArr[i]))
	}
	return res
}

func addByte(b *[]byte, word string, spaceCount int) {
	for i := 0; i < len(word); i++ {
		*b = append(*b, word[i])
	}
	for i := 0; i < spaceCount; i++ {
		*b = append(*b, ' ')
	}
}
