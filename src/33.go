package main

func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0;
		} else {
			return -1;
		}
	}

	left, right := 0, len(nums) - 1;

	if nums[len(nums) - 1] < nums[0] {
		for left <= right {

			mid := (left + right) / 2;

			if nums[mid] == target {
				return mid
			}

			if mid == 0 {
				left = 1;
				right = len(nums) - 1;
				break;
			}

			if nums[mid - 1] < nums[mid] && nums[mid + 1] < nums[mid] {
				if target >= nums[mid + 1] && target <= nums[len(nums) - 1] {
					left = mid + 1;
					right = len(nums) - 1;
					break;
				} else {
					left = 0;
					right = mid - 1;
					break;
				}
			}

			if nums[left] > nums[mid] {
				right = mid
			} else {
				left = mid
			}
		}
	}

	for left <= right {
		mid := (left + right) / 2;
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}