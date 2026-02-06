package main

func isTrionic(nums []int) bool {
	if len(nums) <= 3 {
		return false
	}
	increasing := true
	decreasing := false
	increasing2 := false
	if nums[1] < nums[0] {
		return false
	} 
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1] {
			return false
		}
		if increasing {
			if nums[i - 1] > nums[i] {
				increasing = false
				decreasing = true
				continue;
			}
		}
		if decreasing {
			if nums[i] > nums[i - 1] {
				increasing2 = true
				decreasing = false
			}
		}
		if increasing2 {
			if nums[i - 1] > nums[i] {
				return false
			} 
		}
	}
	return increasing2
}