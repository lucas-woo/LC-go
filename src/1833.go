package main

import "sort"

func maxIceCream(costs []int, coins int) int {
  sort.Ints(costs)
	i := 0
	for i < len(costs) && coins >= costs[i] {
		coins -= costs[i]
		i++
	}
	return i
}