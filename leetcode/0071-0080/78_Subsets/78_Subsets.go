package main

func main() {
	subsets([]int{1, 2, 3})
}

func test() {

}

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}

	t := [][]int{}
	for i := 0; i < len(nums); i++ {
		t = append(t, []int{i})
	}
	for len(t) != 0 {
		res = append(res, t[0])
		for i := t[0][len(t[0])-1] + 1; i < len(nums); i++ {
			temp := make([]int, len(t[0]))
			copy(temp, t[0])
			temp = append(temp, i)
			t = append(t, temp)
		}
		t = t[1:]
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = nums[res[i][j]]
		}
	}

	return res
}
