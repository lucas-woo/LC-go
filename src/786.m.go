package main

import "sort"
type El struct {
	val float64
	i int
	j int
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	sorted := make([]*El, 0)

	for j := len(arr) - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			sorted = append(sorted, &El{
				val: float64(arr[i])/float64(arr[j]),
				i: i,
				j: j,
			})
		}
	}
	sort.Slice(sorted, func (i, j int) bool {
		return sorted[i].val < sorted[j].val
	})
	return []int{sorted[k - 1].i, sorted[k - 1].j}
}