package main

func champagneTower(poured int, query_row int, query_glass int) float64 {

	tower := make([][]float64, 0)

	for i := 1; i <= 100; i++ {
		tower = append(tower, make([]float64, i))
	}
	tower[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {

		for j := 0; j < i + 1; j++ {
			diff := (tower[i][j] - 1)/ 2
			if diff > 0 {
				tower[i + 1][j] += diff
				tower[i + 1][j + 1] += diff
			}
		}
	}
	return min(tower[query_row][query_glass], 1.0)
}	

