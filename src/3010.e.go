package main

func minimumCost(nums []int) int {
	sum := nums[0]
	first := min(nums[1], nums[2])
	second := max(nums[1], nums[2]);

	for i := 3; i < len(nums); i++ {

		if nums[i] < first {
			first, second = nums[i], first
		} else if nums[i] < second {
			second = nums[i]
		}
	}
	return sum + first + second
} 