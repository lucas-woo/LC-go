package main

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	biggest := 1
	for i := 0; i < len(arr); i++ {
		if arr[i] > biggest {
				biggest++
		}
	}
	return min(biggest, len(arr))
}