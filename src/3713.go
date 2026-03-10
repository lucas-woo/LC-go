package main

func longestBalancedddd(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		m := make(map[byte]int, 0)
		numberInMap := 0;
		highestCount := 0;
		numberEqual := 0
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; !ok {
				numberInMap++;
				m[s[j]] = 0;
			}
			m[s[j]]++
			if m[s[j]] == highestCount {
				numberEqual++;
			}
			if m[s[j]] > highestCount {
				highestCount = m[s[j]]
				numberEqual = 1;
			}
			if numberEqual == numberInMap {
				longest = max(longest, j - i + 1)
			}
		}
	}
	return longest
}
