package main

func lengthOfLongestSubstring(s string) int {
	var charHash map[byte]int = make(map[byte]int);
	var maximum int = 0;
	var left int = 0;

	for i, v := range s {
		index, ok := charHash[byte(v)]
		if ok && index >= left {
			left = index + 1;
		}
		maximum = max(maximum, i - left + 1)
		charHash[byte(v)] = i
	}
	return maximum
}
