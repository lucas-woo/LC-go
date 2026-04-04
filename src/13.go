
func romanToInt(s string) int {
	var hashmap map[rune]int = make(map[rune]int);
	hashmap[73] = 1;
	hashmap[86] = 5
	hashmap[88] = 10
	hashmap[76] = 50
	hashmap[67] = 100
	hashmap[68] = 500
	hashmap[77] = 1000

	var currentCount int = 0;
	var currentRune rune;
	var totalCount int = 0
	for i, v := range s {
		if i == 0 {
			currentRune = v
		}
		if v != currentRune {
			if hashmap[v] > hashmap[currentRune] {
				totalCount -= currentCount;
				currentCount = 0;
			} else {
				totalCount += currentCount
				currentCount = 0;
			}
			currentRune = v
		} 
		currentCount += hashmap[v]
	}
	totalCount += currentCount 
	return totalCount
}