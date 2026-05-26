package main 

import "unicode"

func numberOfSpecialChars(word string) int {
 
	m := make(map[byte]bool, 0)
	for i := 0; i < len(word); i++ {
		m[word[i]] = true
	}
	count := 0
	for k := range m {
		lower := byte(unicode.ToLower(rune(k)))
		if lower == byte(k) {
			continue
		}
		if _, ok := m[lower]; ok {
			count++;
		}
	}
	return count
}