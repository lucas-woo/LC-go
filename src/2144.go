package main

import "sort"
func minimumCost(cost []int) int {
	sort.Ints(cost)
	total := 0
	r := 1
	for i := len(cost) - 1; i >= 0; i-- {
		if r % 3 == 0 {
			r = 1
			continue
		}
		r++
		total += cost[i]
	}
	return total
}