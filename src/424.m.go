package main

func characterReplacement(s string, k int) int {
  if k >= len(s) {
		return k
	}
	var contentMap map[byte]int = make(map[byte]int, 0)
	var startMap map[byte]int = make(map[byte]int, 0)

	biggest := k + 1

	for i := 0; i < len(s); i++ {
		curLetter := s[i]
		if _, ok := startMap[curLetter]; !ok {
			startMap[curLetter] = i;
			contentMap[curLetter] = 1;
			continue;
		}
		contentMap[curLetter]++;

		if (i - startMap[v]) - contentMap[v] >




	}
	return biggest
}