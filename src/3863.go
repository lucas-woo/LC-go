package main

import "sort"
func minOperations(s string) int {

	if len(s) == 1 {
		return 0
	}
	if len(s) == 2 {
		if s[0] <= s[1] {
			return 0
		} else {
			return -1
		}
	}
	arr1 := []byte(s)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	if string(arr1) == s {
		return 0
	}
	arr2 := []byte(s[1:])
	arr3 := []byte(s[:len(s) - 1])
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	if string(arr1) == (string(arr1[0]) + string(arr2)) {
		return 1
	}
	sort.Slice(arr3, func(i, j int) bool {
		return arr3[i] < arr3[j]
	})
	if string(arr1) == (string(arr3) + string(arr1[len(s) - 1])) {
		return 1
	}	
	arr3 = append(arr3, s[len(s) - 1])
	arr4 := []byte(arr3[1:])
	sort.Slice(arr4, func(i, j int) bool {
		return arr4[i] < arr4[j]
	})
	if string(arr1) == (string(arr3[0]) + string(arr4)) {
		return 2
	}
	arr2 = []byte((string(s[0]) + string(arr2)))
	arr5 := []byte(arr2[:len(arr2) - 1])
	sort.Slice(arr5, func(i, j int) bool {
		return arr5[i] < arr5[j]
	})

	if string(arr1) == (string(arr5) + string(arr2[len(arr2) - 1])) {
		return 2
	}	
	return 3
}