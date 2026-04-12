package main

func searchInsert(nums []int, target int) int {
  var left int = 0;
	var right int = len(nums) - 1
	for {
		mid := (left + right )/ 2;
		if (left > right) {
			if target < nums[mid] {
				return mid
			} else {
				return mid + 1
			}
		}
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}