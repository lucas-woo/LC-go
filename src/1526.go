package main

func minNumberOperations(target []int) int {
	var retInt int = 0;
	var prev int = 0
	for _, v := range target {
		if v > prev {
			retInt += (v - prev)
		}
		prev = v
	}
	return retInt
}