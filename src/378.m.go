package main
import "sort"
func kthSmallest(matrix [][]int, k int) int {

	sorted := make([]int , len(matrix) * len(matrix))
	p := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			sorted[p] = matrix[i][j]
			p++
		}
	}
	sort.Ints(sorted)
	return sorted[k - 1]
}
