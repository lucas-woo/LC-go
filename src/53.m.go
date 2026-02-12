package main


func maxSubArray(nums []int) int {
	biggest := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + sum < nums[i] {
			sum = nums[i]
			biggest = max(biggest, sum)
			continue
		}
		sum += nums[i]
		biggest = max(biggest, sum)
	}
	return biggest
}