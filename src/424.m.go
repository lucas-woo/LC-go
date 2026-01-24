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

		if (i - startMap[curLetter] + 1) - contentMap[curLetter] > k {

			for j := startMap[curLetter] + 1; j < len(s); j++ {

				if s[j] == s[i] {
					startMap[s[i]] = j;
					contentMap[curLetter]--;
					break;
				}

			}

		} else {
			temp := min((i+1) + (k - (i+1 - contentMap[curLetter])), len(s))
			biggest = max(biggest, temp)
		}


	}

	return biggest
}