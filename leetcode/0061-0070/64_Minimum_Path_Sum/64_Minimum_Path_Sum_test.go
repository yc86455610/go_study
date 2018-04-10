package main

import (
	"testing"
)

func Test(t *testing.T) {
	a := [][]int{[]int{1}}
	ret := minPathSum(a)
	if ret != 1 {
		t.Error("Expected 1 got ", ret)
	}
	a = [][]int{[]int{1, 2}}
	ret = minPathSum(a)
	if ret != 3 {
		t.Fatalf("Expacted 3 got %d", ret)
	}
	a = [][]int{[]int{1}, []int{2}}
	ret = minPathSum(a)
	if ret != 3 {
		t.Fatalf("Expacted 3 got %d", ret)
	}
	a = [][]int{[]int{1, 2}, []int{3, 1}}
	ret = minPathSum(a)
	if ret != 4 {
		t.Fatalf("Expacted 4 got %d", ret)
	}
	a = [][]int{[]int{1, 2, 100}, []int{3, 100, 1}}
	ret = minPathSum(a)
	if ret != 104 {
		t.Fatalf("Expacted 104 got %d", ret)
	}
	a = [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}
	ret = minPathSum(a)
	if ret != 7 {
		t.Fatalf("Expacted 7 got %d", ret)
	}
}
