package main

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func (i, j int) bool {
		return arr[i] < arr[j]
	})
	minDiff := arr[1] - arr[0];
	for i := 1; i < len(arr); i++ {
		minDiff = min(arr[i] - arr[i-1], minDiff)
	}
	var rA [][]int = make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] - arr[i-1] == minDiff {
			rA = append(rA, []int{arr[i-1], arr[i]})
		}
	}
	return rA
}