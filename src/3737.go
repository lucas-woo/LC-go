package main

func countMajoritySubarrays(nums []int, target int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		cnt := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == target {
				cnt++
			}
			if cnt > (j - i + 1 - cnt) {
				total++
			}
		}
	}
	return total
}