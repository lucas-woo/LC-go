func minimumPairRemoval(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	operations := 0;
	
	for i := 0; i < len(nums); i++ {

		var left, right int
		smallest := 1000000000000000000;

		sorted := true
		for j := 0; j < len(nums) - i - 1; j++ {
			if nums[j] > nums[j+1] {
				sorted = false
			}
			if nums[j] + nums[j+1] < smallest {
				left = j
				right = j + 1
				smallest = nums[j] + nums[j+1]
			}
		}
		if sorted {
			break;
		}
		operations++;
		nums[left] = nums[left] + nums[right]
		nums[right] = -1000000000000000000
		for j := 0; j < len(nums) - i - 1; j++ {
			if nums[j] == -1000000000000000000 {
				nums[j] = nums[j+1]
				nums[j+1] = -1000000000000000000
			}
		}
	}
	return operations
}
