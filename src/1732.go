package main

func largestAltitude(gain []int) int {	
	cur := 0;
	biggest := 0

	for i := 0; i < len(gain); i++ {
		cur += gain[i]
		biggest = max(cur, biggest)
	}
	return biggest
}