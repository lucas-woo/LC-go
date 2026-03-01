package main


func minimumPrefixLength(nums []int) int {

	lastBroken := -1;

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i - 1] {
			lastBroken = i - 1
		}
	}

	return lastBroken + 1;

}