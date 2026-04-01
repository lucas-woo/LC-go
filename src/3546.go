package main

func canPartitionGrid(grid [][]int) bool {
	row := make([]int, len(grid))
	col := make([]int, len(grid[0]))
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			row[i] += grid[i][j]
		}
		total += row[i]
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			col[i] += grid[j][i]
		}
	}
	sum := 0
	for i := 0; i < len(row) - 1; i++ {
		sum += row[i];
		if sum == total - sum {
			return true
		}
	}
	sum = 0
	for i := 0; i < len(col) - 1; i++ {
		sum += col[i];
		if sum == total - sum {
			return true
		}
	}	
	return false
}