package main

func reverseBits(n int) int {
	returnNum := 0
	for i := 0; i < 32; i++ {
		returnNum *= 2;
		returnNum += n % 2
		n /= 2
	}
	return returnNum
}