package main

import (
	"fmt"
)

func main() {
	test([]int{0}, []int{0})
	test([]int{1}, []int{1})
	test([]int{2}, []int{2})
	test([]int{1, 0}, []int{0, 1})
	test([]int{0, 2, 1}, []int{0, 1, 2})
	test([]int{0, 1, 2, 0, 2, 1}, []int{0, 0, 1, 1, 2, 2})
}

func test(a []int, b []int) {
	sortColors2(a)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fmt.Println(a, b)
			break
		}
	}
}

func sortColors(nums []int) {
	red := 0
	blue := len(nums) - 1
	for i := 0; i <= blue && red < blue; i++ {
		if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red++
			if nums[i] != 0 {
				i--
			}
		} else if nums[i] == 1 {
			continue
		} else {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
			i--
		}
	}
}

func sortColors2(nums []int) {
	i := 0
	j := 0
	k := 0
	for p := 0; p < len(nums); p++ {
		if nums[p] == 0 {
			nums[k] = 2
			nums[j] = 1
			nums[i] = 0
			k++
			j++
			i++
		} else if nums[p] == 1 {
			nums[k] = 2
			nums[j] = 1
			k++
			j++
		} else if nums[p] == 2 {
			nums[k] = 2
			k++
		}
	}
}
