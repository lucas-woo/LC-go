package main

func constructTransformedArray(nums []int) []int {
	returnArr := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		returnArr[i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {

		if nums[i] < 0 {
			if i + nums[i] < 0 {
				
				if (abss(i + nums[i]) % len(nums)) == 0 {
					returnArr[i] = nums[0]
				} else {
					returnArr[i] = nums[len(nums) - (abss(i + nums[i]) % len(nums))]
				}

			} else {
				returnArr[i] = nums[i + nums[i]]
			}
		} else {
			returnArr[i] = nums[(i + nums[i]) % len(nums)]
		}

	}
	return returnArr
}

func abss(n int) int {
	if n < 0 {
		return -n
	}
	return n
}