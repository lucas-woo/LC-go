package main

import "sort"

func maximumLength(nums []int) int {

	numCountMap := make(map[int]int, 0);
	numMap := make(map[int]int, 0);
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = 1
		if _, ok := numCountMap[nums[i]]; ok {
			numCountMap[nums[i]]++;
		} else {
			numCountMap[nums[i]] = 1
		}
	}
	biggest := 1;
	unique := make([]int, 0)
	sort.Ints(nums);
	if nums[0] != 1 {
			unique = append(unique, nums[0])
	}
	if nums[0] == 1 {
		if numCountMap[1] % 2 == 0 {
			biggest = numCountMap[1] / 2
		} else {
			biggest = (numCountMap[1] + 1) / 2
		}
	}
	prev := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != prev {
			unique = append(unique, nums[i])
			prev = nums[i]
		}
	}

	for i := 0; i < len(unique); i++ {
		if numCountMap[unique[i]] >= 2 {
			if _, ok := numMap[unique[i] * unique[i]]; ok {
				numMap[unique[i] * unique[i]] += numMap[unique[i]];
			}
		} 
		biggest = max(biggest, numMap[unique[i]])
	}
	return (biggest * 2) - 1
}