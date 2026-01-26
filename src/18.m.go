package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var alreadyFound map[[4]int]bool = make(map[[4]int]bool, 0)
	var returnArray [][]int = make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			for right := len(nums) - 1; right > j + 1; right-- {
				left := j + 1;
				for right > left {
					curSum := nums[i] + nums[j] + nums[left] + nums[right]
					if curSum == target {
						temp := [4]int{nums[i], nums[j], nums[left], nums[right]}
						if _, ok := alreadyFound[temp]; ok {
							break;
						}
						alreadyFound[temp] = true;
						
						returnArray = append(returnArray, []int{nums[i], nums[j], nums[left], nums[right]})
						break;
					}
					if curSum < target {
						left++;
					} else {
						right--;
					}
				}				
			}
		}
	}
	return returnArray
}