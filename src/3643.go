package main

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i := y; i < y + k; i++ {
		reverseVertically(grid, x, i, x + k - 1)
	}
	return grid
}

func reverseVertically(grid [][]int, i, j, last int) {
	for i < last {
		grid[i][j], grid[last][j] = grid[last][j], grid[i][j]
		i++;
		last--;
	}
}