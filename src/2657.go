package main

func findThePrefixCommonArray(A []int, B []int) []int {
	seenA := make([]bool, len(A))
	seenB := make([]bool, len(A))
	arr := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		seenA[A[i] - 1] = true;
		seenB[B[i] - 1] = true;
		for j := 0; j < len(A); j++ {
			if seenA[j] &&seenB[j] {
				arr[i]++
			}
		}
	}
	return arr
}