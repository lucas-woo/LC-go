package main

func nextGreaterElements(nums []int) []int {

	stack := make([]int, len(nums))
	ptr := 0
	returnArr := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		for ptr >= 0 && nums[i] > nums[stack[ptr]] {
			returnArr[stack[ptr]] = nums[i]
			ptr--;
		}
		ptr++;
		stack[ptr] = i
	}

	for i := 0; i < len(nums); i++ {
		if i >= stack[ptr] {
			continue
		}
		for ptr >= 0 && nums[i] > nums[stack[ptr]] {
			returnArr[stack[ptr]] = nums[i]
			ptr--;
		}
	}
	for i := 0; i <= ptr; i++ {
		returnArr[stack[i]] = -1
	}
	return returnArr
}