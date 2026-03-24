package main

func constructProductMatrix(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])
	left := make([]int, n*m)
	right := make([]int, n*m)
	flat := make([]int, n*m)
	ind := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			flat[ind] = grid[i][j]
			ind++
		}
	}
	prod := 1
	for i := 0; i < len(flat); i++ {
		left[i] = prod
		prod = (prod * flat[i]) % 12345
	}
	prod = 1
	for i := len(flat) - 1; i >= 0; i-- {
		right[i] = prod;
		prod = (prod * flat[i]) % 12345
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = (left[(i * m) + (j)] * right[(i * m) + (j)]) % 12345
		}
	}
	return grid
}