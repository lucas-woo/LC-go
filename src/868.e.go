package main

func binaryGap(n int) int {

	start := false
	biggest := 0
	count := 0
	for n > 0 {
		if !start {
			if n % 2 == 1 {
				start = true
				count = 1
			}
			n /= 2
		} else {
			if n % 2 == 1 {
				biggest = max(biggest, count)
				start = true
				count = 0
			}
			n /= 2
			count++;
		}
	}
	return biggest
}