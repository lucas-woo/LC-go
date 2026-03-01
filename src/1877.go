package main

import "sort"
func minPairSum(nums []int) int {
	sort.Ints(nums);
	biggest := nums[0] + nums[len(nums) - 1]

	i := 0;
	j := len(nums) - 1;

	for i < j {
		biggest = max(nums[i] + nums[j], biggest)
		i++;
		j--
	}

	return biggest
}