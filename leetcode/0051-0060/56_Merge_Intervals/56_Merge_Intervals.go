package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []Interval{Interval{1, 3}, Interval{2, 6}, Interval{8, 10}, Interval{9, 18}}
	// fmt.Println(merge(a))
	a = []Interval{Interval{1, 4}, Interval{0, 2}, Interval{3, 5}}
	fmt.Println(merge(a))
}

// Interval  Definition for an interval.
type Interval struct {
	Start int
	End   int
}

type sortInterval []Interval

func (s sortInterval) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}
func (s sortInterval) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortInterval) Len() int {
	return len(s)
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	sort.Sort(sortInterval(intervals))
	// fmt.Println(intervals)
	res := []Interval{}
	start := 0
	end := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= intervals[end].End {
			if intervals[i].End > intervals[end].End {
				end = i
			}
		} else {
			res = append(res, Interval{intervals[start].Start, intervals[end].End})
			start = i
			end = i
		}
		// fmt.Println(i, start, end, res)
	}
	res = append(res, Interval{intervals[start].Start, intervals[end].End})
	return res
}
