package main

func maximumJumps(nums []int, target int) int {

	biggest := -1;
	m := make(map[[2]int]int, 0)
	maximumJumpsBruteForce(nums, target, 0, 1, 0, &biggest, m)
	if biggest <= 0 {
		return -1
	}
	return biggest
}

func maximumJumpsBruteForce(nums []int, target, start, end int, jumpCount int, biggest *int, m map[[2]int]int) {

	if v, ok := m[[2]int{start,end}]; ok {
		if v >= jumpCount {
			return
		}
	} 
	m[[2]int{start,end}] = jumpCount

	if *biggest >= (len(nums) - end + jumpCount) {
		return
	}
	if start == len(nums) - 1 {
		*biggest = max(*biggest, jumpCount)
		return
	}	
	if end >= len(nums) || start >= len(nums) {
		return 
	}
	t := nums[end] - nums[start] 
	if t >= -target && t <= target { 
		maximumJumpsBruteForce(nums, target, end, end+1, jumpCount+1, biggest,m)
	}
	maximumJumpsBruteForce(nums, target, start, end + 1, jumpCount, biggest,m)
}
