package main

func minElement(nums []int) int {
smallest := 100000000
	for i := 0; i < len(nums); i++ {
		t := 0
		for nums[i] > 0 {
			t += nums[i] % 10
			nums[i] /= 10
		}
		smallest = min(smallest, t)
	}
	return smallest
}