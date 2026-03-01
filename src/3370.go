package main
func smallestNumber(n int) int {
	var comp int = 1;
	var temp int = 1
	for {
		if comp >= n {
			return comp 
		}
		temp *= 2
		comp += temp
	}
}