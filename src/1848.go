package main

func getMinDistance(nums []int, target int, start int) int {
	sm := len(nums) + 100
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			sm = min(sm, abs(i, start))
		}
	}
}

func abs(a,b int) int {
	if a < b {
		return b - a
	} 
	return a - b
}