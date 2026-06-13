package main

import "strings"
func mapWordWeights(words []string, weights []int) string {
	var sb strings.Builder
	for i := 0; i < len(words); i++ {
		w := 0
		for j := 0; j < len(words[i]); j++ {
			w += weights[abs(int(words[i][j]) - 97)]
		}
		w %= 26
		
		sb.WriteByte(byte(abs(122-w)))
	}
	return sb.String()
}

func abs(n int) int {
	if n < 0 {
		return -n
	} 
	return n
}