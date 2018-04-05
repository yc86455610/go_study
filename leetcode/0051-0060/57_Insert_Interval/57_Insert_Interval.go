package main

import "fmt"

func main() {
	a := []Interval{Interval{1, 3}, Interval{6, 9}}
	fmt.Println(insert(a, Interval{2, 5}))
	a = []Interval{Interval{1, 2}, Interval{3, 5}, Interval{6, 7}, Interval{8, 10}, Interval{12, 16}}
	fmt.Println(insert(a, Interval{4, 9}))
	a = []Interval{Interval{1, 5}}
	fmt.Println(insert(a, Interval{2, 7}))
	a = []Interval{Interval{1, 5}, Interval{6, 8}}
	fmt.Println(insert(a, Interval{0, 9}))
}

// Interval  Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	res := []Interval{}

	start := 0
	min := newInterval.Start

	for i := 0; i < len(intervals); i++ {
		if newInterval.Start > intervals[i].End {
			start++
		} else {
			if newInterval.Start > intervals[i].Start {
				min = intervals[i].Start
			}
			break
		}
	}

	end := len(intervals)
	max := newInterval.End

	for i := len(intervals) - 1; i >= 0; i-- {
		if newInterval.End < intervals[i].Start {
			end--
		} else {
			if newInterval.End < intervals[i].End {
				max = intervals[i].End
			}
			break
		}
	}

	for i := 0; i < start; i++ {
		res = append(res, intervals[i])
	}
	res = append(res, Interval{min, max})

	for i := end; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}
