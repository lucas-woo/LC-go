package main

import "sort"

func check(nums []int) bool {
  arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)
	complete := true
	for i := 0; i < len(arr); i++ {
		if arr[i] != nums[i] {
			complete = false
			break
		}
	}
    if complete {
        return true
    }
	start := nums[0]
	i := 1;
	for i < len(nums) {
		if nums[i] < nums[i-1] {
			break;
		}
        
		i++
	}
    if nums[i] > start {
        return false
    }

	for j := i + 1; j < len(nums); j++ {
		if nums[j] < nums[j-1] || nums[j] > start {
			return false
		} 
	}
	return true
}