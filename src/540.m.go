func singleNonDuplicate(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	left := 0;
	right := len(nums) - 1;

	for left <= right {

		mid := (left + right) / 2;
		if mid == 0 || mid == len(nums) - 1 {
			return nums[mid]
		}
		
		if nums[mid - 1] != nums[mid] && nums[mid + 1] != nums[mid] {
			return nums[mid]
		}
		
		if nums[mid - 1] == nums[mid] && (mid + 1) % 2 != 0 {
			right = mid - 1;
		} else if nums[mid - 1] == nums[mid] && (mid + 1) % 2 == 0 {
			left = mid + 1;
		} else if nums[mid + 1] == nums[mid] && (mid + 1) % 2 == 0 {
			right = mid - 1;
		} else {
			left = mid + 1;
		}

	}
	return -1
}