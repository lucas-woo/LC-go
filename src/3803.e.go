package main

func residuePrefixes(s string) int {
	var m map[byte]bool = make(map[byte]bool, 0);
	rN := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			m[s[j]] = true
		}
		if len(m) == ((i+1)%3) {
			rN++
		}
		m = make(map[byte]bool, 0);
	}
	return rN
}