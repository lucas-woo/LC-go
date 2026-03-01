package main

import "sort"
import "math/bits"
type a struct {
	ones int
	Val int
}

func sortByBits(arr []int) []int {
	rA := make([]*a, len(arr))
  for i := 0; i < len(arr); i++ {
		rA[i] = &a{
			ones: int(bits.OnesCount(uint(arr[i]))),
			Val: arr[i],
		}
	}
	sort.Slice(rA, func(i, j int) bool {
		if rA[i].ones == rA[j].ones {
			return rA[i].Val < rA[j].Val
		}
		return rA[i].ones < rA[j].ones
	})
	for i := 0; i < len(arr); i++ {
		arr[i] = rA[i].Val
	}
	return arr
}