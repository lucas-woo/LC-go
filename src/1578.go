package main

func minCost(colors string, neededTime []int) int {
	var maxTime int = 0
	var prev rune;
	var prevIndex int;
	for i, v := range colors {
		if prev == v {
			if neededTime[i] < neededTime[prevIndex] {
				maxTime += neededTime[i]
				} else {
				maxTime += neededTime[prevIndex]
				prevIndex = i
			}
		} else {
			prev = v
			prevIndex = i
		}
	}
	return maxTime
}