package main

import "sort"
func threeSum(nums []int) [][]int {
	var mapArr map[[3]int]bool = make(map[[3]int]bool, 0)
	var returnArray [][]int = make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		left := i + 1;
		right := len(nums) - 1;

		for left < right {
			curSum := nums[i] + nums[left] + nums[right]
			if curSum == 0 {
				curArr := [3]int{nums[i], nums[left], nums[right]}
				if _, ok := mapArr[curArr]; ok {
					left++;
					continue;
				}
				mapArr[curArr] = true
				returnArray = append(returnArray, []int{nums[i], nums[left], nums[right]})
			}
			if curSum < 0 {
				left++;
			} else {
				right--;
			}
		}

		
	}
	return returnArray
}