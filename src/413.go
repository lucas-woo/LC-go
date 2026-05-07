package main

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	count := 0;
	start := 0;
	end := 1;
	diff := nums[1] - nums[0];
	for i := 1; i < len(nums); i++ {
		if nums[i] - nums[i - 1] != diff {
			start = i - 1;
			end = i + 1;
			diff = nums[i] - nums[i - 1];
			continue;
		}
		end++;
		if end - start >= 3 {
			count += (end - start - 2)
		}
	}
	return count
}