package main

func countSubmatrices(grid [][]int, k int) int {
	for i := 0; i < len(grid); i++ {
		sum := 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
			grid[i][j] = sum
		}
	}

	accepted := 0

	for i := len(grid[0]) - 1; i >= 0; i-- {
		prev := 0;
		for j := 0; j < len(grid); j++ {
			prev += grid[j][i]
			if prev > k {
				break;
			}
			accepted++;
		}
	}

	return accepted

}