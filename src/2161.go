package main

func pivotArray(nums []int, pivot int) []int {
	arr := make([]int, 0)
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == pivot {
			p++
		}
		if nums[i] < pivot {
			arr = append(arr, nums[i])
		}
	}
	for i := 0; i < p; i++ {
		arr = append(arr, pivot)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > pivot {
			arr = append(arr, nums[i])
		}
	}
	return arr
}