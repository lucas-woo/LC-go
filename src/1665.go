package main

import "sort"
func minimumEffort(tasks [][]int) int {


	sort.Slice(tasks, func(i,j int) bool {
		diffi := tasks[i][1] - tasks[i][0]
		diffj := tasks[j][1] - tasks[j][0]
		if diffi == diffj {
			return tasks[i][1] < tasks[j][0]
		}
		return diffi < diffj
	})

	total := 0

	cur := 0

	for i := len(tasks) - 1; i >= 0; i-- {
		if cur < tasks[i][1] {
			temp := tasks[i][1] - cur
			cur += temp
			total += temp
		}
		cur -= tasks[i][0]
	}
	return total
}
