package main

func maxTotalValue(nums []int, k int) int64 {
	smallest := nums[0]
	biggest := nums[0]

	for i := 0; i < len(nums); i++ {
		smallest = min(smallest, nums[i])
		biggest = max(biggest, nums[i])
	}
	return int64(biggest-smallest) * int64(k)
}