package main

func beautifulArray(n int) []int {
	if n == 1 {
		return []int{1}
	}
	more := beautifulArray((n/2) + (n % 2))
	less := beautifulArray(n/2)

	var rA []int = make([]int, n);
	for i := 0; i < len(more); i++ {
		rA[i] = (more[i])*2-1
	}
	for i := 0; i < len(less); i++ {
		rA[len(more) + i] = less[i]*2
	}
	return rA
}