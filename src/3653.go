package main

func xorAfterQueries(nums []int, queries [][]int) (r int) {
    t := 1000000007
	for i := 0; i < len(queries); i++ {
		a, b, c, d := queries[i][0], queries[i][1], queries[i][2], queries[i][3]
		for j := a; j <= b; j += c {
			nums[j] = (nums[j] * d) % t
		}
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			r = nums[i]
		} else {
			r ^= nums[i]
		}
	}
	return
}