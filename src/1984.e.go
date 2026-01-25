package main

import "sort"

func minimumDifference(nums []int, k int) int {
  sort.Ints(nums)
	smallest := nums[len(nums) - 1] - nums[0] 

	for i := 0; i <= len(nums) - k; i++ {

		smallest = min(nums[i+k-1] - nums[i], smallest)
	}
	return smallest
}