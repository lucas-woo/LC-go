package main

func isPalindrome(x int) bool {
	copy := x
	rev := 0
	for {
		if copy <= 0 {
			break
		}
		rev *= 10
		rev += (copy % 10)
		copy /= 10
	}
	return rev == x;
}