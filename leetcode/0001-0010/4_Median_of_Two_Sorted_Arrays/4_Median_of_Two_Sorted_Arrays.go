package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 = []int{1, 3, 5}
	var num2 = []int{2}

	fmt.Println(num1)
	fmt.Println(num2)

	fmt.Println(findMedianSortedArrays(num1, num2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	l1 := len(nums1)
	l2 := len(nums2)

	head := 0
	tail := l1
	var c1, c2, L1, R1, L2, R2 int
	for head <= tail {
		c1 = (head + tail) / 2
		c2 = (l1+l2)/2 - c1

		if c1 == 0 {
			L1 = math.MinInt64
		} else {
			L1 = nums1[c1-1]
		}
		if c2 == 0 {
			L2 = math.MinInt64
		} else {
			L2 = nums2[c2-1]
		}
		if c1 == l1 {
			R1 = math.MaxInt64
		} else {
			R1 = nums1[c1]
		}
		if c2 == l2 {
			R2 = math.MaxInt64
		} else {
			R2 = nums2[c2]
		}

		if L1 > R2 {
			tail = c1 - 1
		} else if R1 < L2 {
			head = c1 + 1
		} else {
			break
		}
	}
	if (l1+l2)%2 == 0 {
		return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2.0
	}
	return math.Min(float64(R1), float64(R2))
}
