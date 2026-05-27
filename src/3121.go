package main 

import "unicode"

func numberOfSpecialChars(word string) int {
 
	m := make(map[byte]int, 0)
	first := make(map[byte]int, 0)
	for i := 0; i < len(word); i++ {
		m[word[i]] = i
		if _, ok := first[word[i]]; !ok {
			first[word[i]] = i
		}
	}
	count := 0
	for k, v := range m {
		upper := byte(unicode.ToUpper(rune(k)))
		if upper == byte(k) {
			continue
		}
		if f, ok := first[upper]; ok && v < f {
			count++
		}
	}
	return count
}