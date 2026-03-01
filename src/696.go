package main
func countBinarySubstrings(s string) int {
	ones := 0
	zero := 0
	prev := s[0]
	total := 0

	for i := 0; i < len(s); i++ {
		if s[i] == prev {
			if s[i] == '0' {
				zero++;
			} else {
				ones++;
			} 
			continue
		}
		total += min(ones, zero)		
		if s[i] == '0' {
			zero = 0;
		} else {
			ones = 0;
		}
		prev = s[i]
		if s[i] == prev {
			if s[i] == '0' {
				zero++;
			} else {
				ones++;
			} 
			continue
		}		
	}
	total += min(ones, zero)
	return total
}
