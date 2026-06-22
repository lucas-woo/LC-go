package main


func maxNumberOfBalloons(text string) int {
	m := make(map[byte]int, 0)
	m['b'] = 0;
	m['a'] = 0
	m['l'] = 0
	m['o'] = 0
	m['n'] = 0
	for i := 0; i < len(text); i++ {
		if _, ok := m[text[i]]; ok {
			m[text[i]]++;
		}
	}
	count := len(text)
	for b, v := range m {
		if b == 'l' || b == 'o' {
			count = min(count, v/2)
		} else {
			count = min(count, v)
		}
	}
	return count
}
